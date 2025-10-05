package service

import (
	"context"

	innerpb "github.com/tgkzz/ecom/innerpb/auth"
)

type AuthGrpc interface {
	Register(ctx context.Context, request *innerpb.RegisterRequest) (response *innerpb.RegisterResponse, err error)
}

type Service struct {
	Auth AuthGrpc
}

func NewService(authClient innerpb.AuthClient) *Service {
	return &Service{
		Auth: NewAuth(authClient),
	}
}
