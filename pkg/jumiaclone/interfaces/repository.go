package interfaces

import (
	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/domain/dao"
)

type Repository interface {
	CreateUser(user *dao.User) (*dao.User, error)
	SaveOTP(otp *dao.OTPPayload) error
}
