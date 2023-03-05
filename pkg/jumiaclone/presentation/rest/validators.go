package rest

import (
	"fmt"

	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/application"
	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/domain/dao"
	"github.com/asaskevich/govalidator"
)

func validateCreateUserInput(input *dao.UserProfile) (*dao.UserProfile, error) {
	if input == nil {
		return nil, fmt.Errorf("nil payload provided")
	}
	if input.FirstName == "" || input.LastName == "" || input.PhoneNumber == "" || input.PassWord == "" || input.Email == "" {
		return nil, fmt.Errorf("invalid request data, ensure you have these fields: `first_name`, `last_name`, `phone_number`, `password")
	}

	email := govalidator.IsEmail(input.Email)
	if !email {
		return nil, fmt.Errorf(
			"invalid request data, invalid email address",
		)
	}
	normalizedPhoneNumber, err := application.NormalizeMSISDN(
		input.PhoneNumber,
	)
	if err != nil {
		return nil, err
	}
	input.PhoneNumber = *normalizedPhoneNumber

	return input, nil
}

func validateOtpPayload(otp *dao.OTPPayload) (*dao.OTPPayload, error) {
	if otp == nil {
		return nil, fmt.Errorf("nil payload provided")
	}
	if otp.PhoneNumber == "" {
		return nil, fmt.Errorf("invalid request data, ensure you have the phone_number field")
	}
	if otp.OTPPassword == "" {
		return nil, fmt.Errorf("invalid request data, ensure you have the otp_password field")
	}
	normalizedPhoneNumber, err := application.NormalizeMSISDN(otp.PhoneNumber)
	if err != nil {
		return nil, err
	}
	otp.PhoneNumber = *normalizedPhoneNumber
	return otp, nil
}
