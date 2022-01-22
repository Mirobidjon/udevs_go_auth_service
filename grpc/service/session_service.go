package service

import (
	"context"
	"errors"
	"fmt"
	"time"
	"upm/udevs_go_auth_service/config"
	"upm/udevs_go_auth_service/grpc/client"
	"upm/udevs_go_auth_service/pkg/logger"
	"upm/udevs_go_auth_service/pkg/security"
	"upm/udevs_go_auth_service/storage"

	pb "upm/udevs_go_auth_service/genproto/auth_service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type sessionService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	pb.UnimplementedSessionServiceServer
}

func NewSessionService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *sessionService {
	return &sessionService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (s *sessionService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	res := &pb.LoginResponse{}

	if len(req.Username) < 6 {
		err := errors.New("invalid username")
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if len(req.Password) < 6 {
		err := errors.New("invalid password")
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := s.strg.User().GetByUsername(req.Username)
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	match, err := security.ComparePassword(user.Password, req.Password)
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if !match {
		err := errors.New("username or password is wrong")
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if user.Active < 0 {
		err := errors.New("user is not active")
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if user.Active == 0 {
		err := errors.New("user hasn't been activated yet")
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	fmt.Println(user.ExpiresAt)

	expiresAt, err := time.Parse(config.DatabaseTimeLayout, user.ExpiresAt)
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if expiresAt.Unix() < time.Now().Unix() {
		err := errors.New("user has been expired")
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	res.UserFound = true
	res.User = user

	clientType, err := s.strg.ClientType().GetByPK(&pb.ClientTypePrimaryKey{
		Id: user.ClientTypeId,
	})
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	res.ClientType = clientType

	clientPlatform, err := s.strg.ClientPlatform().GetByPK(&pb.ClientPlatformPrimaryKey{
		Id: user.ClientPlatformId,
	})
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	res.ClientPlatform = clientPlatform

	client, err := s.strg.Client().GetByPK(&pb.ClientPrimaryKey{
		ClientPlatformId: user.ClientPlatformId,
		ClientTypeId:     user.ClientTypeId,
	})
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if client.LoginStrategy != pb.LoginStrategies_STANDARD {
		err := errors.New("incorrect login strategy")
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	role, err := s.strg.Role().GetByPK(&pb.RolePrimaryKey{Id: user.RoleId})
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	res.Role = role

	// TODO - Delete all old sessions & refresh token has this function too
	rowsAffected, err := s.strg.Session().DeleteExpiredUserSessions(user.Id)
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	s.log.Info("Login--->DeleteExpiredUserSessions", logger.Any("rowsAffected", rowsAffected))

	userSessionList, err := s.strg.Session().GetSessionListByUserID(user.Id)
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	res.Sessions = userSessionList.Sessions

	sessionPKey, err := s.strg.Session().Create(&pb.CreateSessionRequest{
		ProjectId:        user.ProjectId,
		ClientPlatformId: user.ClientPlatformId,
		ClientTypeId:     user.ClientTypeId,
		UserId:           user.Id,
		RoleId:           user.RoleId,
		Ip:               "0.0.0.0",
		Data:             "additional json data",
		ExpiresAt:        time.Now().Add(config.RefreshTokenExpiresInTime).Format(config.DatabaseTimeLayout),
	})
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	session, err := s.strg.Session().GetByPK(sessionPKey)
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// TODO - wrap in a function
	m := map[string]interface{}{
		"id":                 session.Id,
		"project_id":         session.ProjectId,
		"client_platform_id": session.ClientPlatformId,
		"client_type_id":     session.ClientTypeId,
		"user_id":            session.UserId,
		"role_id":            session.RoleId,
		"ip":                 session.Data,
		"data":               session.Data,
	}

	accessToken, err := security.GenerateJWT(m, config.AccessTokenExpiresInTime, s.cfg.SecretKey)
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	refreshToken, err := security.GenerateJWT(m, config.RefreshTokenExpiresInTime, s.cfg.SecretKey)
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	res.Token = &pb.Token{
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		CreatedAt:        session.CreatedAt,
		UpdatedAt:        session.UpdatedAt,
		ExpiresAt:        session.ExpiresAt,
		RefreshInSeconds: int32(config.AccessTokenExpiresInTime.Seconds()),
	}

	return res, nil
}

func (s *sessionService) Logout(ctx context.Context, req *pb.LogoutRequest) (*emptypb.Empty, error) {
	claims, err := security.ExtractClaims(req.AccessToken, s.cfg.SecretKey)
	if err != nil {
		s.log.Error("!!!Logout--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// TODO - wrap in a function
	id := claims["id"].(string)
	// projectID := claims["project_id"].(string)
	// clientPlatformID := claims["client_platform_id"].(string)
	// clientTypeID := claims["client_type_id"].(string)
	// userID := claims["user_id"].(string)
	// roleID := claims["role_id"].(string)
	// ip := claims["ip"].(string)
	// data := claims["data"].(string)

	rowsAffected, err := s.strg.Session().Delete(&pb.SessionPrimaryKey{Id: id})
	if err != nil {
		s.log.Error("!!!Logout--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	s.log.Info("---Logout--->", logger.Any("claims", claims))
	s.log.Info("---Logout--->", logger.Any("rowsAffected", rowsAffected))

	return &emptypb.Empty{}, nil
}

// func (s *sessionService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
// 	return
// }

// func (s *sessionService) HasAccess(ctx context.Context, req *pb.HasAccessRequest) (*pb.HasAccessResponse, error) {
// 	return
// }