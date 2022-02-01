package handlers

import (
	"context"
	"upm/udevs_go_auth_service/api/http"

	"upm/udevs_go_auth_service/genproto/auth_service"

	"upm/udevs_go_auth_service/pkg/util"

	"github.com/gin-gonic/gin"
)

// AddRole godoc
// @ID create_role
// @Router /role [POST]
// @Summary Create Role
// @Description Create Role
// @Tags Role
// @Accept json
// @Produce json
// @Param role body auth_service.AddRoleRequest true "AddRoleRequestBody"
// @Success 201 {object} http.Response{data=auth_service.Role} "Role data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) AddRole(c *gin.Context) {
	var role auth_service.AddRoleRequest

	err := c.ShouldBindJSON(&role)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PermissionService().AddRole(
		context.Background(),
		&role,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// UpdateRole godoc
// @ID update_role
// @Router /role [PUT]
// @Summary Update Role
// @Description Update Role
// @Tags Role
// @Accept json
// @Produce json
// @Param role body auth_service.UpdateRoleRequest true "UpdateRoleRequestBody"
// @Success 200 {object} http.Response{data=auth_service.Role} "Role data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateRole(c *gin.Context) {
	var role auth_service.UpdateRoleRequest

	err := c.ShouldBindJSON(&role)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PermissionService().UpdateRole(
		context.Background(),
		&role,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// RemoveRole godoc
// @ID delete_role
// @Router /role/{role-id} [DELETE]
// @Summary Delete Role
// @Description Get Role
// @Tags Role
// @Accept json
// @Produce json
// @Param role-id path string true "role-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) RemoveRole(c *gin.Context) {
	roleID := c.Param("role-id")

	if !util.IsValidUUID(roleID) {
		h.handleResponse(c, http.InvalidArgument, "role id is an invalid uuid")
		return
	}

	resp, err := h.services.PermissionService().RemoveRole(
		context.Background(),
		&auth_service.RolePrimaryKey{
			Id: roleID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}

// CreatePermission godoc
// @ID create_permission
// @Router /permission [POST]
// @Summary Create Permission
// @Description Create Permission
// @Tags Permission
// @Accept json
// @Produce json
// @Param permission body auth_service.CreatePermissionRequest true "CreatePermissionRequestBody"
// @Success 201 {object} http.Response{data=auth_service.Permission} "Permission data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreatePermission(c *gin.Context) {
	var permission auth_service.CreatePermissionRequest

	err := c.ShouldBindJSON(&permission)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PermissionService().CreatePermission(
		context.Background(),
		&permission,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetPermissionList godoc
// @ID get_permission_list
// @Router /permission [GET]
// @Summary Get Permission List
// @Description  Get Permission List
// @Tags Permission
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=auth_service.GetPermissionListResponse} "GetPermissionListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetPermissionList(c *gin.Context) {
	offset, err := h.getOffsetParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.services.PermissionService().GetPermissionList(
		context.Background(),
		&auth_service.GetPermissionListRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetPermissionByID godoc
// @ID get_permission_by_id
// @Router /permission/{permission-id} [GET]
// @Summary Get Permission By ID
// @Description Get Permission By ID
// @Tags Permission
// @Accept json
// @Produce json
// @Param permission-id path string true "permission-id"
// @Success 200 {object} http.Response{data=auth_service.Permission} "PermissionBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetPermissionByID(c *gin.Context) {
	permissionID := c.Param("permission-id")

	if !util.IsValidUUID(permissionID) {
		h.handleResponse(c, http.InvalidArgument, "permission id is an invalid uuid")
		return
	}

	resp, err := h.services.PermissionService().GetPermissionByID(
		context.Background(),
		&auth_service.PermissionPrimaryKey{
			Id: permissionID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdatePermission godoc
// @ID update_permission
// @Router /permission [PUT]
// @Summary Update Permission
// @Description Update Permission
// @Tags Permission
// @Accept json
// @Produce json
// @Param permission body auth_service.UpdatePermissionRequest true "UpdatePermissionRequestBody"
// @Success 200 {object} http.Response{data=auth_service.Permission} "Permission data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdatePermission(c *gin.Context) {
	var permission auth_service.UpdatePermissionRequest

	err := c.ShouldBindJSON(&permission)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PermissionService().UpdatePermission(
		context.Background(),
		&permission,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeletePermission godoc
// @ID delete_permission
// @Router /permission/{permission-id} [DELETE]
// @Summary Delete Permission
// @Description Get Permission
// @Tags Permission
// @Accept json
// @Produce json
// @Param permission-id path string true "permission-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeletePermission(c *gin.Context) {
	permissionID := c.Param("permission-id")

	if !util.IsValidUUID(permissionID) {
		h.handleResponse(c, http.InvalidArgument, "permission id is an invalid uuid")
		return
	}

	resp, err := h.services.PermissionService().DeletePermission(
		context.Background(),
		&auth_service.PermissionPrimaryKey{
			Id: permissionID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}

// UpsertScope godoc
// @ID upsert_scope
// @Router /upsert-scope [POST]
// @Summary Upsert Scope
// @Description Upsert Scope
// @Tags UpsertScope
// @Accept json
// @Produce json
// @Param upsert-scope body auth_service.UpsertScopeRequest true "UpsertScopeRequestBody"
// @Success 201 {object} http.Response{data=auth_service.Role} "Role data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpsertScope(c *gin.Context) {
	var upsert_scope auth_service.UpsertScopeRequest

	err := c.ShouldBindJSON(&upsert_scope)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PermissionService().UpsertScope(
		context.Background(),
		&upsert_scope,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// AddPermissionScope godoc
// @ID add_permission_scope
// @Router /permission-scope [POST]
// @Summary Create PermissionScope
// @Description Create PermissionScope
// @Tags PermissionScope
// @Accept json
// @Produce json
// @Param permission-scope body auth_service.AddPermissionScopeRequest true "AddPermissionScopeRequestBody"
// @Success 201 {object} http.Response{data=auth_service.PermissionScope} "PermissionScope data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) AddPermissionScope(c *gin.Context) {
	var permission_scope auth_service.AddPermissionScopeRequest

	err := c.ShouldBindJSON(&permission_scope)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PermissionService().AddPermissionScope(
		context.Background(),
		&permission_scope,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// RemovePermissionScope godoc
// @ID delete_permission_scope
// @Router /permission-scope [DELETE]
// @Summary Delete PermissionScope
// @Description Get PermissionScope
// @Tags PermissionScope
// @Accept json
// @Produce json
// @Param permission-scope body auth_service.PermissionScopePrimaryKey true "PermissionScopePrimaryKeyBody"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) RemovePermissionScope(c *gin.Context) {
	var permission_scope auth_service.PermissionScopePrimaryKey

	err := c.ShouldBindJSON(&permission_scope)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PermissionService().RemovePermissionScope(
		context.Background(),
		&permission_scope,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}

// AddRolePermission godoc
// @ID add_role_permission
// @Router /role-permission [POST]
// @Summary Create RolePermission
// @Description Create RolePermission
// @Tags RolePermission
// @Accept json
// @Produce json
// @Param role-permission body auth_service.AddRolePermissionRequest true "AddRolePermissionRequestBody"
// @Success 201 {object} http.Response{data=auth_service.RolePermission} "RolePermission data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) AddRolePermission(c *gin.Context) {
	var role_permission auth_service.AddRolePermissionRequest

	err := c.ShouldBindJSON(&role_permission)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PermissionService().AddRolePermission(
		context.Background(),
		&role_permission,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// RemoveRolePermission godoc
// @ID delete_role_permission
// @Router /role-permission [DELETE]
// @Summary Delete RolePermission
// @Description Get RolePermission
// @Tags RolePermission
// @Accept json
// @Produce json
// @Param role-permission body auth_service.RolePermissionPrimaryKey true "RolePermissionPrimaryKeyBody"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) RemoveRolePermission(c *gin.Context) {
	var role_permission auth_service.RolePermissionPrimaryKey

	err := c.ShouldBindJSON(&role_permission)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PermissionService().RemoveRolePermission(
		context.Background(),
		&role_permission,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}