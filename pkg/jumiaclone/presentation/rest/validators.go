package rest

import (
	"fmt"

	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/domain/dao"
)

func validateCreateUserInput(input *dao.User) (*dao.User, error) {
	if input == nil {
		return nil, fmt.Errorf("nil payload provided")
	}
	if input.FirstName == "" || input.LastName == "" || input.PhoneNumber == "" || input.PassWord == "" {
		return nil, fmt.Errorf("invalid request data, ensure you have these fields: `first_name`, `last_name`, `phone_number`, `password")
	}
	return input, nil
}
