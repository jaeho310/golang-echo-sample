package service

import (
	"errors"
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

	// TODO 1. 중복된 이름이 들어왔을때 svc단에서 처리
	isDuplicated, err := userServiceImpl.UserRepository.DuplicatedCheck(user.Name)
	// 중복된경우
	if isDuplicated {
		return nil, errors.New("중복된 유저입니다.")
	} else if err != nil {
		// 중복회원조회할때 DB에러
		return nil, err
	}

	// TODO 2. 중복된 이름이 들어왔을때 repo단에서 처리
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
