package services

import "github.com/Israel-Ferreira/todo-s3/pkg/models"

type UserServiceImpl struct {
}

func (usi UserServiceImpl) GetById(id uint64) (*models.User, error) {
	return nil, nil
}

func (usi UserServiceImpl) GetAll() ([]*models.User, error) {
	return nil, nil
}

func (usi UserServiceImpl) DeleteById(id uint64) error {
	return nil
}

func (usi UserServiceImpl) UploadPhoto()
