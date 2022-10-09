package repositories

import "github.com/Israel-Ferreira/todo-s3/pkg/models"

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindById(uint64) (*models.User, error)
	DeleteById(uint64) error
	Update(uint64, models.User) error
	UpdateProfileImageUrl(uint64, string) error
	Create(models.User) (uint64, error)
}
