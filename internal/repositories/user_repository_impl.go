package repositories

import (
	"database/sql"
	"log"

	"github.com/Israel-Ferreira/todo-s3/internal/db"
	"github.com/Israel-Ferreira/todo-s3/pkg/models"
	"github.com/Israel-Ferreira/todo-s3/pkg/repositories"
)

type UserRepositoryImpl struct{}

func (usr *UserRepositoryImpl) FindAll() ([]models.User, error) {
	conn, err := db.OpenDbConnection()

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer conn.Close()

	query, err := conn.Query("select u.id, u.username, u.email, u.profile_photo_url from users u")

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer query.Close()

	users := []models.User{}

	for query.Next() {
		user := models.NewUser()

		var imgUrl sql.NullString

		if err := query.Scan(&user.ID, &user.Name, &user.Email, &imgUrl); err != nil {
			return nil, err
		}

		user.UserProfileImgUrl = imgUrl.String

		users = append(users, *user)
	}

	return users, nil
}

func (usr *UserRepositoryImpl) FindById(id uint64) (*models.User, error) {
	conn, err := db.OpenDbConnection()

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	row := conn.QueryRow("select u.id, u.username, u.email, u.profile_photo_url from users where u.id = $1", id)

	user := models.NewUser()

	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.UserProfileImgUrl); err != nil {
		return nil, err
	}

	return nil, nil
}

func (usr *UserRepositoryImpl) Create(user models.User) (uint64, error) {

	conn, err := db.OpenDbConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	stmt, err := conn.Prepare(
		`insert into users(username, email, user_password) values ($1, $2, $3)`,
	)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&user.Name, &user.Email, &user.Password)

	if err != nil {
		return 0, err
	}

	return 1, nil
}

func (usr *UserRepositoryImpl) DeleteById(id uint64) error {
	conn, err := db.OpenDbConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	stmt, err := conn.Prepare("delete from users where id = $1")

	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err = stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func (usr *UserRepositoryImpl) UpdateProfileImageUrl(id uint64, url string) error {

	conn, err := db.OpenDbConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	stmt, err := conn.Prepare("update users set profile_photo_url = $1 where id = $2")

	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err = stmt.Exec(url, id); err != nil {
		return err
	}

	return nil
}

func (usr *UserRepositoryImpl) Update(id uint64, user models.User) error {
	return nil
}

func (usr *UserRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	conn, err := db.OpenDbConnection()

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	query := conn.QueryRow("select u.id, u.username, u.email, u.profile_photo_url from users u where u.email = $1", email)

	user := models.NewUser()

	if err := query.Scan(&user.ID, &user.Name, &user.Email, &user.UserProfileImgUrl); err != nil {
		return nil, err
	}

	return user, nil

}

func NewUserRepository() repositories.UserRepository {
	return &UserRepositoryImpl{}
}
