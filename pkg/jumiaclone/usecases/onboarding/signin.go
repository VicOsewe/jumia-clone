package onboarding

import (
	"fmt"

	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/application"
	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/domain/dao"
)

func (o *Onboarding) SignInByEmail(email, password string) (*dao.UserProfile, error) {
	user, err := o.repository.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user: %v", err)
	}
	if user.ID == "" {
		return nil, fmt.Errorf("user with this email does not exist")
	}
	if !user.IsVerified {
		return nil, fmt.Errorf("user is not verified please verify this user before proceeding to sign in")
	}

	matched := application.ComparePIN(password, user.Salt, user.PassWord, nil)

	if !matched {
		return nil, fmt.Errorf("wrong password provided, please enter the correct password")
	}

	//TODO: give user access
	return user, nil
}

func (o *Onboarding) SingInByPhoneNumber(phoneNumber, password string) (*dao.UserProfile, error) {
	normalizedPhoneNumber, err := application.NormalizeMSISDN(
		phoneNumber,
	)
	if err != nil {
		return nil, err
	}
	user, err := o.repository.GetUserByPhoneNumber(*normalizedPhoneNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user: %v", err)
	}
	if user.ID == "" {
		return nil, fmt.Errorf("user with this email does not exist")
	}
	if !user.IsVerified {
		return nil, fmt.Errorf("user is not verified please verify this user before proceeding to sign in")
	}

	matched := application.ComparePIN(password, user.Salt, user.PassWord, nil)
	if !matched {
		return nil, fmt.Errorf("wrong password provided, please enter the correct password")
	}

	//TODO: give user access
	return user, nil
}
