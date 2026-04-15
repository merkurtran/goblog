package user

import "github.com/merkurtran/goblog/app/models"

type User struct {
	models.BaseModel

	Name     string
	Email    string
	Password string
}
