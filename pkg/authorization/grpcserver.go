package authorization

import (
	"authorization/pkg/api"
	"context"
	"log"

	password "github.com/vzglad-smerti/password_hash"
)

type GRPCServer struct {
	api.UnimplementedAuthorizationServer
}

func (s *GRPCServer) Register(ctx context.Context, req *api.UserMeta) (*api.SuccessfulRegister, error) {
	hash, err := password.Hash(req.GetPassword())
	if err != nil {
		log.Fatal(err)
	}

	return &api.SuccessfulRegister{Login: req.GetLogin(),
		Password: hash}, nil
}
