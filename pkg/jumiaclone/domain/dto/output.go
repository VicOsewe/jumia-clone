package dto

import "github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/domain/dao"

type UserResponse struct {
	Profile dao.UserProfile
}

// APIResponseMessage represents the response for generic happy cases for the RESTFUL apis
type APIResponseMessage struct {
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Body       interface{} `json:"body"`
}
