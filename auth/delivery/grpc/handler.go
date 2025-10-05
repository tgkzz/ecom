package grpc

import (
	"context"
	"log/slog"

	"github.com/tgkzz/ecom/auth/service"
	innerpb "github.com/tgkzz/ecom/innerpb/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TODO: twilio revovery code DRCXA3R68ECMJ3T318SUD43P save telegram

type authRouter struct {
	logger  *slog.Logger
	service *service.Service
	innerpb.UnimplementedAuthServer
}

func (a *authRouter) Register(ctx context.Context, req *innerpb.RegisterRequest) (*innerpb.RegisterResponse, error) {
	const op = "authRouter.Register"

	log := a.logger.WithGroup(op)

	// TODO: send otp to number using https://www.twilio.com/en-us. use test service

	if err := a.service.OtpSend.SendOtp(req.GetPhoneNum(), req.GetPassword()); err != nil {
		log.Error("error while sending otp", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	// TODO: write in database new user as in register step

	// TODO: save otp for 3 * time.Minute

}
