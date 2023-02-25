package onboarding

import (
	"fmt"
	"log"
	"time"

	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/application"
	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/domain/dao"
	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/interfaces"
	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/interfaces/services"

	"github.com/pkg/errors"
	"github.com/pquerna/otp/totp"
)

const (
	otpMessage = "%s is your Jumia verification codeS"
)

// OnboardingUsecase is a representation of the usecase's contract
type OnboardingUsecase interface {
	CreateUser(user *dao.User) (*dao.User, error)
}

// Onboarding sets up the onboarding's usecase layer with all the necessary dependencies
type Onboarding struct {
	totpOpts   totp.GenerateOpts
	repository interfaces.Repository
	sms        services.SMS
}

// NewOnboarding initializes the Onboarding usecase instance that meets all the preconditions checks
func NewOnboarding(repo interfaces.Repository, sms services.SMS) *Onboarding {
	on := &Onboarding{
		totpOpts: totp.GenerateOpts{
			Issuer:      application.MustGetEnvVar("ISSUER"),
			AccountName: application.MustGetEnvVar("ACCOUNTNAME"),
		},
		repository: repo,
		sms:        sms,
	}
	on.checkPreConditions()
	return on
}

func (o *Onboarding) checkPreConditions() {
	if o.totpOpts.AccountName == "" {
		log.Panicf("error, account name  must be provided")
	}
	if o.totpOpts.Issuer == "" {
		log.Panicf("error, issuer must be provided")
	}
}

func (o *Onboarding) CreateUser(user *dao.User) (*dao.User, error) {

	_, encryptedPassword := application.EncryptPIN(user.PassWord, nil)

	user.Verified = false
	user.PassWord = encryptedPassword
	userResponse, err := o.repository.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)

	}

	otpCode, err := o.GenerateOTP()
	if err != nil {
		return nil, fmt.Errorf("failed to generate otp code: %v", err)
	}

	otp := dao.OTPPayload{
		PhoneNumber: user.PhoneNumber,
		OTPPassword: otpCode,
		IsValid:     true,
		Timestamp:   time.Now(),
	}

	err = o.repository.SaveOTP(&otp)
	if err != nil {
		return nil, fmt.Errorf("failed to save otp: %v", err)
	}
	msg := fmt.Sprintf(otpMessage, otpCode)

	_, err = o.sms.SendSMS(msg, user.PhoneNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to send ot sms: %v", err)
	}
	return userResponse, nil
}

func (s *Onboarding) GenerateOTP() (string, error) {
	key, err := totp.Generate(s.totpOpts)
	if err != nil {
		return "", errors.Wrap(err, "generateOTP")
	}
	code, err := totp.GenerateCode(key.Secret(), time.Now())
	if err != nil {
		return "", errors.Wrap(err, "generateOTP > GenerateCode")
	}
	return code, nil
}
