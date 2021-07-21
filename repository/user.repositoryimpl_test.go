package repository_test

import (
	"github.com/stretchr/testify/assert"
	"platform-sample/infrastructure/database"
	"platform-sample/infrastructure/server"
	"platform-sample/model"
	"platform-sample/repository"
	"testing"
)

func initRepository() *repository.UserRepositoryImpl {
	mockDb := database.SqlStore{}.GetMockDb()
	mockServer := server.Server{MainDb: mockDb}
	return mockServer.InjectUserRepository()
}

func Test_CreateAndDeleteUser(t *testing.T) {
	// given
	newUser := &model.User{}
	newUser.Name = "Kitty"

	// when
	userRepositoryImpl := initRepository()
	user, err := userRepositoryImpl.Save(newUser)
	if err != nil {
		t.Error(err)
	}
	// then
	assert.Equal(t, user, newUser)

	// 실제 영속성을 건드므로 유저를 지우는걸 해야한다.
	err = userRepositoryImpl.DeleteById(int(user.ID))
	if err != nil {
		t.Error(err)
	}
}

func Test_Duplicated_User(t *testing.T) {
	userRepositoryImpl := initRepository()
	newUser := &model.User{}
	newUser.Name = "Kitty"

	_, err1 := userRepositoryImpl.Save(newUser)
	if err1 != nil {
		t.Error(err1)
	}

	_, err2 := userRepositoryImpl.Save(newUser)
	assert.Error(t, err2)
}
