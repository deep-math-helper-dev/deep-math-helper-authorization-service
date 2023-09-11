package authorization

import (
	"authorization/internal/db"
	"authorization/internal/entity"
	"authorization/pkg/api"
	"context"
)

type GRPCServer struct {
	api.UnimplementedAuthorizationServer
}

func (s *GRPCServer) Register(ctx context.Context, req *api.UserMeta) (*api.SuccessfulRegister, error) {
	login, hash := entity.TakeData(req.GetLogin(), req.GetPassword()) // formatting users data
	db.Add(login, hash)                                               // add data into db
	return &api.SuccessfulRegister{Login: login, Password: hash}, nil // server response
}
