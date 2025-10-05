package service

import (
	"log/slog"

	"github.com/tgkzz/ecom/auth/repository"
)

type OtpSender interface {
	SendOtp(otp, phoneNum string) error
}

type OtpGenerator interface {
	Generate() (string, error)
}

type Service struct {
	OtpSend     OtpSender
	OtpGenerate OtpGenerator
}

func NewService(repo *repository.Repository, logger *slog.Logger) *Service {
	return &Service{
		OtpSend:     NewOtp(repo, logger),
		OtpGenerate: NewOtp(repo, logger),
	}
}
