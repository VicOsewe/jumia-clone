package onboarding

import (
	"time"

	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/domain/dto"
	"github.com/pkg/errors"
	"github.com/pquerna/otp/totp"
)

// OnboardingUsecase is a representation of the usecase's contract
type OnboardingUsecase interface {
}

// Onboarding sets up the onboarding's usecase layer with all the necessary dependencies
type Onboarding struct {
	totpOpts totp.GenerateOpts
}

// NewOnboarding initializes the Onboarding usecase instance that meets all the preconditions checks
func NewOnboarding() *Onboarding {
	on := &Onboarding{}
	on.checkPreConditions()
	return on
}

func (o *Onboarding) checkPreConditions() {

}

func (o *Onboarding) CreateUser() (*dto.User, error) {
	//check if all required fields are present
	// create a record of the user in the db and mark it as a non-verified user
	// generate and send a verification code to the user
	return nil, nil
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
