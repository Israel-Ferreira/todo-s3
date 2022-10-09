package services

import (
	"bytes"
	"mime/multipart"
	"net/http"

	"github.com/Israel-Ferreira/todo-s3/internal/config"
	"github.com/Israel-Ferreira/todo-s3/pkg/data"
	"github.com/Israel-Ferreira/todo-s3/pkg/models"
	"github.com/Israel-Ferreira/todo-s3/pkg/repositories"
	"github.com/Israel-Ferreira/todo-s3/pkg/services"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/private/protocol/rest"
	"github.com/aws/aws-sdk-go/service/s3"
)

type UserServiceImpl struct {
	repository repositories.UserRepository
}

func (usi UserServiceImpl) GetById(id uint64) (*models.User, error) {
	user, err := usi.repository.FindById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (usi UserServiceImpl) GetAll() ([]models.User, error) {
	users, err := usi.repository.FindAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (usi UserServiceImpl) DeleteById(id uint64) error {

	if err := usi.repository.DeleteById(id); err != nil {
		return err
	}

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

	s3Client := s3.New(config.S3Session)

	_, err := s3Client.PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(config.ConfigVars.AwsBucketName),
		Key:                aws.String(key),
		Body:               bytes.NewReader(buffer),
		ContentLength:      aws.Int64(int64(size)),
		ACL:                aws.String("public-read"),
		ContentType:        aws.String(http.DetectContentType(buffer)),
		ContentDisposition: aws.String("attachment"),
	})

	if err != nil {
		return "", err
	}

	req, _ := s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(config.ConfigVars.AwsBucketName),
		Key:    aws.String(key),
	})

	rest.Build(req)

	fileUrl := req.HTTPResponse.Request.URL.String()

	if err := usi.repository.UpdateProfileImageUrl(userId, fileUrl); err != nil {
		return "", err
	}

	return fileUrl, nil
}

func NewUserServiceImpl(repository repositories.UserRepository) services.UserService {
	return &UserServiceImpl{repository: repository}
}
