package services

import "github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/domain/dto"

type SMS interface {
	SendSMS(message, to string) (*dto.SendMessageResponse, error)
}
