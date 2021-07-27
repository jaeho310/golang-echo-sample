package repository

import (
	"platform-sample/model"
)

type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindById(id int) (*model.User, error)
	DeleteById(id int) error
	Save(*model.User) (*model.User, error)
	Update(*model.User) (*model.User, error)
}
