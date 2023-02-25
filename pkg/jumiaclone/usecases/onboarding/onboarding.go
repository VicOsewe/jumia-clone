package onboarding

import "github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/domain/dto"

// OnboardingUsecase is a representation of the usecase's contract
type OnboardingUsecase interface {
}

// Onboarding sets up the onboarding's usecase layer with all the necessary dependencies
type Onboarding struct {
}

// NewOnboarding initializes the Onboarding usecase instance that meets all the preconditions checks
func NewOnboarding() *Onboarding {
	o := &Onboarding{}
	o.checkPreConditions()
	return o
}

func (o *Onboarding) checkPreConditions() {

}

func (o *Onboarding) CreateUser() (*dto.User, error) {
	//check if all required fields are present
	// create a record of the user in the db and mark it as a non-verified user
	// generate and send a verification code to the user
	return nil, nil
}
