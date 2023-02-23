package dao

import "time"

type OTPPayload struct {
	PhoneNumber string    `json:"phone_number"`
	OTPPassword string    `json:"otp_password"`
	Timestamp   time.Time `json:"timestamp,omitempty" firestore:"timestamp"`
	IsValid     bool      `json:"is_valid"`
}
