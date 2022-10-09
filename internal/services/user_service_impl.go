package services

import (
	"bytes"
	"errors"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/Israel-Ferreira/todo-s3/internal/config"
	"github.com/Israel-Ferreira/todo-s3/pkg/data"
	"github.com/Israel-Ferreira/todo-s3/pkg/models"
	"github.com/Israel-Ferreira/todo-s3/pkg/repositories"
	"github.com/Israel-Ferreira/todo-s3/pkg/security"
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
		log.Println(err)
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

	if userData.Name == "" {
		return 0, errors.New("o nome não pode estar em branco")
	}

	if userData.Email == "" {
		return 0, errors.New("o email não pode estar em branco")
	}

	if userData.Password == "" {
		return 0, errors.New("a senha não pode estar em branco")
	}

	_, err := usi.repository.FindByEmail(userData.Email)

	if err == nil {
		return 0, errors.New("email já cadastrado")
	}

	hashPass, err := security.HashPassword(userData.Password)

	if err != nil {
		return 0, err
	}

	user := models.User{
		Name:     userData.Name,
		Email:    userData.Email,
		Password: string(hashPass),
	}

	insertedId, err := usi.repository.Create(user)

	if err != nil {
		return 0, err
	}

	return insertedId, nil
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
