package interfaces

import (
	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/domain/dao"
)

type Repository interface {
	CreateUser(user *dao.UserProfile) (*dao.UserProfile, error)
	GetUserByPhoneNumber(phoneNumber string) (*dao.UserProfile, error)
	GetUserByEmail(email string) (*dao.UserProfile, error)
	UpdateUser(user *dao.UserProfile) (*dao.UserProfile, error)
	SaveOTP(otp *dao.OTPPayload) error
	GetOTP(phoneNumber, otp string) (*dao.OTPPayload, error)
	UpdateOTP(otp *dao.OTPPayload) (*dao.OTPPayload, error)
}
