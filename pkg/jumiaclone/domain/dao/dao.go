package dao

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OTPPayload struct {
	PhoneNumber string    `json:"phone_number"`
	OTPPassword string    `json:"otp_password"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
	IsValid     bool      `json:"is_valid"`
}

type User struct {
	ID          string `json:"uid,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	MiddleName  string `json:"middle_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email,omitempty"`
	Verified    bool   `json:"verified,omitempty"`
	PassWord    string `json:"password,omitempty"`
}

// BeforeCreate User hook ensures that before a new session is created, a new unique UUID
// is added
func (s *User) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New().String()
	return
}
