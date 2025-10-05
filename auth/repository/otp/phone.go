package otp

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log/slog"

	innerpb "github.com/tgkzz/ecom/innerpb/auth"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type optStatus struct{}

const phoneFrom = "+12173938365"

type PhoneOtp struct {
	twilioClient *twilio.RestClient
	logger       *slog.Logger
}

func NewPhoneOtp(accountSid, authToken string, logger *slog.Logger) Otp {
	return &PhoneOtp{
		twilioClient: twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: accountSid,
			Password: authToken,
		}),
		logger: logger,
	}
}

func (p *PhoneOtp) GenerateOtp() (string, error) {
	const length = 6

	buf := make([]byte, length)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}

	otpCharsLen := len(otpChars)
	for i := 0; i < otpCharsLen; i++ {
		buf[i] = otpChars[int(buf[i])%otpCharsLen]
	}

	return string(buf), nil
}

func (p *PhoneOtp) SendOtp(otp, phoneNum string) error {
	const op = "PhoneOtp.SendOtp"

	var params twilioApi.CreateMessageParams

	params.SetTo(phoneNum)
	params.SetBody(fmt.Sprintf("your otp code is %s", otp))
	params.SetFrom(phoneFrom)

	message, err := p.twilioClient.Api.CreateMessage(&params)
	if err != nil {
		return err
	}

	if message == nil {
		return errors.New(op + " twilio message is nil")
	}

	status := *message.Status
	if message.Status != nil && status == innerpb.OtpStatus_QUEUED.String() {
		return nil
	}

	return errors.New(op + " twilio message is invalid")
}
