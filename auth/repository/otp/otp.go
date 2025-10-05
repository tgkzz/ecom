package otp

const otpChars = "1234567890"

type Otp interface {
	SendOtp(otp, phoneNum string) error
	GenerateOtp() (string, error)
}
