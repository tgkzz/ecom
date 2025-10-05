package service

import (
	"context"
	"log/slog"

	innerpb "github.com/tgkzz/ecom/innerpb/auth"
)

type auth struct {
	authClient innerpb.AuthClient
	logger     *slog.Logger
}

func NewAuth(authClient innerpb.AuthClient) AuthGrpc {
	return &auth{
		authClient: authClient,
	}
}

func (a *auth) Register(ctx context.Context, request *innerpb.RegisterRequest) (response *innerpb.RegisterResponse, err error) {
	const op = "auth.Register"

	log := a.logger.WithGroup(op)

	response, err = a.authClient.Register(ctx, request)
	if err != nil {
		log.Error("error in client register", "op", op, "error", err)
		return
	}

	return
}
