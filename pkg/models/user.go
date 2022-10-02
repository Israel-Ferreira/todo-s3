package models

type User struct {
	ID                uint64 `json:"id"`
	Name              string `json:"name"`
	BirthDate         string `json:"birth_date"`
	UserProfileImgUrl string `json:"user_profile_img_url"`
	Email             string `json:"email"`
	Password          string `json:"-"`
}

func NewUser() *User {
	return &User{}
}
