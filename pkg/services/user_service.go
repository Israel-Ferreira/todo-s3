package services

import (
	"github.com/Israel-Ferreira/todo-s3/pkg/data"
	"github.com/Israel-Ferreira/todo-s3/pkg/models"
)

type UserService interface {
	GetAll() ([]*models.User, error)
	GetById(uint64) (*models.User, error)
	Create(data.UserData) (uint64, error)
	Update(uint64, data.UserData) error
	DeleteById(uint64) error
	UploadPhoto(uint64) (string, error)
}
