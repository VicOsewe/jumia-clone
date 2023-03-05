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
	otpMessage = "%s is your Jumia verification code"
)

// OnboardingUsecase is a representation of the usecase's contract
type OnboardingUsecase interface {
	CreateUser(user *dao.UserProfile) (*dao.UserProfile, error)
	VerifyPhoneNumber(otp *dao.OTPPayload) (bool, error)
	CheckIfEmailExists(email string) (*dao.UserProfile, error)
	CheckIfPhoneNumberExists(phoneNumber string) (*dao.UserProfile, error)
	SignInByEmail(email, password string) (*dao.UserProfile, error)
	SingInByPhoneNumber(phoneNumber, password string) (*dao.UserProfile, error)
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
			AccountName: application.MustGetEnvVar("ACCOUNT_NAME"),
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

func (o *Onboarding) CreateUser(user *dao.UserProfile) (*dao.UserProfile, error) {

	us, err := o.repository.GetUserByPhoneNumber(user.PhoneNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user by phone number: %v", err)

	}
	if len(us.ID) != 0 {
		return nil, fmt.Errorf("phone number is already in use by another user")
	}

	use, err := o.repository.GetUserByEmail(user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user by email: %v", err)

	}

	if len(use.ID) != 0 {
		return nil, fmt.Errorf("phone number is already in use by another user")
	}

	salt, encryptedPassword := application.EncryptPIN(user.PassWord, nil)

	user.IsVerified = false
	user.PassWord = encryptedPassword
	user.Salt = salt
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

func (o *Onboarding) GenerateOTP() (string, error) {
	key, err := totp.Generate(o.totpOpts)
	if err != nil {
		return "", errors.Wrap(err, "generateOTP")
	}
	code, err := totp.GenerateCode(key.Secret(), time.Now())
	if err != nil {
		return "", errors.Wrap(err, "generateOTP > GenerateCode")
	}
	return code, nil
}

func (o *Onboarding) CheckIfPhoneNumberExists(phoneNumber string) (*dao.UserProfile, error) {
	user, err := o.repository.GetUserByPhoneNumber(phoneNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user by phone number: %v", err)
	}
	return user, nil
}

func (o *Onboarding) CheckIfEmailExists(email string) (*dao.UserProfile, error) {
	user, err := o.repository.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user by email: %v", err)
	}
	return user, nil
}

func (o *Onboarding) VerifyPhoneNumber(otp *dao.OTPPayload) (bool, error) {
	// TODO:ensure user exists????
	otpPayload, err := o.repository.GetOTP(otp.PhoneNumber, otp.OTPPassword)
	if err != nil {
		return false, fmt.Errorf("failed to get otp from database: %v", err)
	}
	if len(otpPayload.OTPPassword) == 0 {
		return false, fmt.Errorf("no matching verification codes found")
	}
	if !otpPayload.IsValid {
		return false, fmt.Errorf("verification code is not valid")
	}
	user := dao.UserProfile{
		PhoneNumber: otp.PhoneNumber,
		IsVerified:  true,
	}

	_, err = o.repository.UpdateUser(&user)
	if err != nil {
		return false, err
	}

	otpPay := dao.OTPPayload{
		PhoneNumber: otp.PhoneNumber,
		IsValid:     false,
	}
	_, err = o.repository.UpdateOTP(&otpPay)
	if err != nil {
		return false, err
	}

	return true, nil
}
