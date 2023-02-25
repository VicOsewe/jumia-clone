package repository

import "github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/domain/dto"

type Repository interface {
	CreateUser(user *dto.User) (*dto.User, error)
}
