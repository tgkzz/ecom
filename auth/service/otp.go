package service

import (
	"log/slog"

	"github.com/tgkzz/ecom/auth/repository"
)

type Otp struct {
	repos  *repository.Repository
	logger *slog.Logger
}

func NewOtp(repo *repository.Repository, logger *slog.Logger) *Otp {
	return &Otp{
		repos:  repo,
		logger: logger,
	}
}

func (otp *Otp) SendOtp(otpCode string, phoneNum string) error {
	const op = "otp.SendOtp"

	log := otp.logger.With("op", op, "phoneNum", phoneNum)

	if err := otp.repos.Otp.SendOtp(otpCode, phoneNum); err != nil {
		log.Error("failed to send otp code", "err", err)
	}

	return nil
}

func (otp *Otp) Generate() (string, error) {
	const op = "otp.GenerateOtp"

	log := otp.logger.With("op", op)

	res, err := otp.repos.Otp.GenerateOtp()
	if err != nil {
		log.Error("failed to generate otp", "err", err)
		return "", err
	}

	return res, nil
}
