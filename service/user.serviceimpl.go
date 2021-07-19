package service

import (
	"platform-sample/model"
	"platform-sample/repository"
)

type UserServiceImpl struct {
	repository.UserRepository
}

func (UserServiceImpl) NewUserServiceImpl(repository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repository}
}

func (userServiceImpl *UserServiceImpl) GetUsers() ([]*model.User, error) {
	return userServiceImpl.UserRepository.FindAll()
}

func (userServiceImpl *UserServiceImpl) CreateUser(user *model.User) (*model.User, error) {
	return userServiceImpl.UserRepository.Save(user)
}

func (userServiceImpl *UserServiceImpl) DeleteUser(id int) error {
	return userServiceImpl.UserRepository.DeleteById(id)
}

func (userServiceImpl *UserServiceImpl) GetUser(id int) (*model.User, error) {
	return userServiceImpl.UserRepository.FindById(id)
}

func (userServiceImpl *UserServiceImpl) UpdateUser(user *model.User) (*model.User, error) {
	return userServiceImpl.UserRepository.Update(user)
}
