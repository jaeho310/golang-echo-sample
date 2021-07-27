package repository

import (
	"platform-sample/model"
)

type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindById(id int) (*model.User, error)
	DeleteById(id int) error
	Save(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
}
