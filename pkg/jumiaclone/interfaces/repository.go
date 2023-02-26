package interfaces

import (
	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/domain/dao"
)

type Repository interface {
	CreateUser(user *dao.User) (*dao.User, error)
	GetUserByPhoneNumber(phoneNumber string) (*dao.User, error)
	GetUserByEmail(email string) (*dao.User, error)
	SaveOTP(otp *dao.OTPPayload) error
}
