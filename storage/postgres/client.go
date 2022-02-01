package postgres

import (
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/pkg/util"
	"upm/udevs_go_auth_service/storage"

	"github.com/jmoiron/sqlx"
)

type clientRepo struct {
	db *sqlx.DB
}

func NewClientRepo(db *sqlx.DB) storage.ClientRepoI {
	return &clientRepo{
		db: db,
	}
}

func (r *clientRepo) Add(projectID string, entity *pb.AddClientRequest) (err error) {
	query := `INSERT INTO "client" (
		project_id,
		client_platform_id,
		client_type_id,
		login_strategy
	) VALUES (
		$1,
		$2,
		$3,
		$4
	)`

	_, err = r.db.Exec(query,
		projectID,
		entity.ClientPlatformId,
		entity.ClientTypeId,
		entity.LoginStrategy.String(),
	)

	return err
}

func (r *clientRepo) GetByPK(pKey *pb.ClientPrimaryKey) (res *pb.Client, err error) {
	res = &pb.Client{}
	query := `SELECT
		project_id,
		client_platform_id,
		client_type_id,
		login_strategy
	FROM
		"client"
	WHERE
		client_platform_id = $1 AND client_type_id = $2`

	row, err := r.db.Query(query, pKey.ClientPlatformId, pKey.ClientTypeId)
	if err != nil {
		return res, err
	}
	defer row.Close()

	if row.Next() {
		var loginStrategy string

		err = row.Scan(
			&res.ProjectId,
			&res.ClientPlatformId,
			&res.ClientTypeId,
			&loginStrategy,
		)

		res.LoginStrategy = pb.LoginStrategies(pb.LoginStrategies_value[loginStrategy])

		if err != nil {
			return res, err
		}
	} else {
		return res, storage.ErrorNotFound
	}

	return res, nil
}

func (r *clientRepo) Update(entity *pb.UpdateClientRequest) (rowsAffected int64, err error) {
	query := `UPDATE "client" SET
		login_strategy = :login_strategy,
		updated_at = now()
	WHERE
		client_platform_id = :client_platform_id AND client_type_id = :client_type_id`

	params := map[string]interface{}{
		"client_platform_id": entity.ClientPlatformId,
		"client_type_id":     entity.ClientTypeId,
		"login_strategy":     entity.LoginStrategy.String(),
	}

	result, err := r.db.NamedExec(query, params)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}

func (r *clientRepo) Remove(pKey *pb.ClientPrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM "client" WHERE client_platform_id = $1 AND client_type_id = $2`

	result, err := r.db.Exec(query, pKey.ClientPlatformId, pKey.ClientTypeId)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}

func (r *clientRepo) GetList(queryParam *pb.GetClientListRequest) (res *pb.GetClientListResponse, err error) {
	res = &pb.GetClientListResponse{}
	params := make(map[string]interface{})
	query := `SELECT
		project_id,
		client_platform_id,
		client_type_id,
		login_strategy
	FROM
		"client"`
	filter := " WHERE 1=1"
	order := " ORDER BY created_at"
	arrangement := " DESC"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if util.IsValidUUID(queryParam.ProjectId) {
		params["project_id"] = queryParam.ProjectId
		filter += " AND (project_id = :project_id)"
	}

	if len(queryParam.Search) > 0 {
		params["search"] = queryParam.Search
		filter += " AND (login_strategy ILIKE '%' || :search || '%')"
	}

	if queryParam.Offset > 0 {
		params["offset"] = queryParam.Offset
		offset = " OFFSET :offset"
	}

	if queryParam.Limit > 0 {
		params["limit"] = queryParam.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM "client"` + filter
	row, err := r.db.NamedQuery(cQ, params)
	if err != nil {
		return res, err
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(
			&res.Count,
		)
		if err != nil {
			return res, err
		}
	}

	q := query + filter + order + arrangement + offset + limit
	rows, err := r.db.NamedQuery(q, params)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &pb.Client{}
		var loginStrategy string
		err = rows.Scan(
			&obj.ProjectId,
			&obj.ClientPlatformId,
			&obj.ClientTypeId,
			&loginStrategy,
		)
		if err != nil {
			return res, err
		}
		obj.LoginStrategy = pb.LoginStrategies(pb.LoginStrategies_value[loginStrategy])
		res.Clients = append(res.Clients, obj)
	}

	return res, nil
}

func (r *clientRepo) GetMatrix(req *pb.GetClientMatrixRequest) (res *pb.GetClientMatrixResponse, err error) {
	if !util.IsValidUUID(req.ProjectId) {
		return res, storage.ErrorProjectId
	}

	queryClientPlatform := `SELECT
		id,
		project_id,
		name,
		subdomain
	FROM
		"client_platform"
	WHERE
		project_id = $1`

	clientPlatformRows, err := r.db.Query(queryClientPlatform, req.ProjectId)
	if err != nil {
		return res, err
	}
	defer clientPlatformRows.Close()

	for clientPlatformRows.Next() {
		obj := &pb.ClientPlatform{}
		err = clientPlatformRows.Scan(
			&obj.Id,
			&obj.ProjectId,
			&obj.Name,
			&obj.Subdomain,
		)
		if err != nil {
			return res, err
		}
		res.ClientPlatforms = append(res.ClientPlatforms, obj)
	}

	queryClientType := `SELECT
		id,
		project_id,
		name,
		confirm_by,
		self_register,
		self_recover
	FROM
		"client_type"
	WHERE
		project_id = $1`

	clientTypeRows, err := r.db.Query(queryClientType, req.ProjectId)
	if err != nil {
		return res, err
	}
	defer clientTypeRows.Close()

	for clientTypeRows.Next() {
		obj := &pb.ClientType{}
		var confirmBy string
		err = clientTypeRows.Scan(
			&obj.Id,
			&obj.ProjectId,
			&obj.Name,
			&confirmBy,
			&obj.SelfRegister,
			&obj.SelfRecover,
		)
		if err != nil {
			return res, err
		}
		obj.ConfirmBy = pb.ConfirmStrategies(pb.ConfirmStrategies_value[confirmBy])
		res.ClientTypes = append(res.ClientTypes, obj)
	}

	queryClient := `SELECT
		project_id,
		client_platform_id,
		client_type_id,
		login_strategy
	FROM
		"client"
	WHERE
		project_id = $1`

	clientRows, err := r.db.Query(queryClient, req.ProjectId)
	if err != nil {
		return res, err
	}
	defer clientRows.Close()

	for clientRows.Next() {
		obj := &pb.Client{}
		var loginStrategy string
		err = clientRows.Scan(
			&obj.ProjectId,
			&obj.ClientPlatformId,
			&obj.ClientTypeId,
			&loginStrategy,
		)
		if err != nil {
			return res, err
		}
		obj.LoginStrategy = pb.LoginStrategies(pb.LoginStrategies_value[loginStrategy])
		res.Clients = append(res.Clients, obj)
	}

	return
}
