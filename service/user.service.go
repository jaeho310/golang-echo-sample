package service

import "platform-sample/model"

type UserService interface {
	GetUsers() ([]*model.User, error)
	CreateUser(*model.User) (*model.User, error)
	DeleteUser(int) error
	GetUser(int) (*model.User, error)
	UpdateUser(*model.User) (*model.User, error)
}
