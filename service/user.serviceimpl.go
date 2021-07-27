package service

import (
	"platform-sample/model"
	"platform-sample/repository"
)

type UserServiceImpl struct {
	repository.UserRepository
	repository.CardRepository
}

func (UserServiceImpl) NewUserServiceImpl(userRepository repository.UserRepository,
	cardRepository repository.CardRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepository, cardRepository}
}

func (userServiceImpl *UserServiceImpl) GetUsers() ([]*model.User, error) {
	return userServiceImpl.UserRepository.FindAll()
}

func (userServiceImpl *UserServiceImpl) CreateUser(user *model.User) (*model.User, error) {
	return userServiceImpl.UserRepository.Save(user)
}

func (userServiceImpl *UserServiceImpl) DeleteUser(id int) error {
	err := userServiceImpl.UserRepository.DeleteById(id)
	if err != nil {
		return err
	}
	//err = userServiceImpl.CardRepository.DeleteByUserId(id)
	//if err != nil {
	//	return err
	//}

	return nil
}

func (userServiceImpl *UserServiceImpl) GetUser(id int) (*model.User, error) {
	return userServiceImpl.UserRepository.FindById(id)
}

func (userServiceImpl *UserServiceImpl) UpdateUser(user *model.User) (*model.User, error) {
	return userServiceImpl.UserRepository.Update(user)
}
