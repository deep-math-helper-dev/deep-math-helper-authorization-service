package authorization

import (
	"authorization/pkg/api"
	"context"
)

type GRPCServer struct {
	api.UnimplementedAuthorizationServer
}

func (s *GRPCServer) Register(ctx context.Context, req *api.UserMeta) (*api.SuccessfulRegister, error) {
	return &api.SuccessfulRegister{Login: req.GetLogin(),
		Password: req.GetPassword()}, nil
}
