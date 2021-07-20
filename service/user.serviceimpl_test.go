package service_test

import (
	"github.com/stretchr/testify/assert"
	mocks "platform-sample/mocks/repository"
	"platform-sample/model"
	"platform-sample/service"
	"testing"
)

func TestGetUser(t *testing.T) {
	mockRepository := &mocks.UserRepository{}
	userServiceImpl := service.UserServiceImpl{}.NewUserServiceImpl(mockRepository)

	// given
	mockUser := &model.User{Name: "Tom"}
	mockRepository.On("FindById", 1).Return(mockUser, nil)

	// when
	user, err := userServiceImpl.GetUser(1)
	if err != nil {
		t.Error(err)
	}

	// then
	assert.Equal(t, user, mockUser)
}
