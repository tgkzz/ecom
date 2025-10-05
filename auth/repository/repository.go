package repository

import (
	"log/slog"

	"github.com/tgkzz/ecom/auth/repository/otp"
)

type Repository struct {
	Otp otp.Otp
}

func NewRepository(twillioAccountSID, twillioAccountToken string, logger *slog.Logger) *Repository {
	return &Repository{
		Otp: otp.NewPhoneOtp(twillioAccountSID, twillioAccountToken, logger),
	}
}
