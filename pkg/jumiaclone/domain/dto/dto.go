package dto

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          string `json:"uid,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	MiddleName  string `json:"middle_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email,omitempty"`
	Verified    bool   `json:"verified,omitempty"`
	PassWord    string `json:"pass_word,omitempty"`
}

// BeforeCreate User hook ensures that before a new session is created, a new unique UUID
// is added
func (s *User) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New().String()
	return
}
