package services

import (
	"bytes"
	"mime/multipart"
	"net/http"

	"github.com/Israel-Ferreira/todo-s3/internal/config"
	"github.com/Israel-Ferreira/todo-s3/pkg/data"
	"github.com/Israel-Ferreira/todo-s3/pkg/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

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

func (usi UserServiceImpl) Create(userData data.UserData) (uint64, error) {
	return 1, nil
}

func (usi UserServiceImpl) Update(id uint64, userData data.UserData) error {
	return nil
}

func (usi UserServiceImpl) UploadPhoto(userId uint64, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {

	size := fileHeader.Size
	buffer := make([]byte, size)

	file.Read(buffer)

	key := "profile-photos/" + fileHeader.Filename

	_, err := s3.New(config.S3Session).PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(config.ConfigVars.AwsBucketName),
		Key:                aws.String(key),
		Body:               bytes.NewReader(buffer),
		ContentLength:      aws.Int64(int64(size)),
		ACL: aws.String("public-read"),
		ContentType:        aws.String(http.DetectContentType(buffer)),
		ContentDisposition: aws.String("attachment"),
	})

	if err != nil {
		return "", err
	}

	return "", nil
}
