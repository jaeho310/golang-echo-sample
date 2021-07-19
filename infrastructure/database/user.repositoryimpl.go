package database

import (
	"github.com/jinzhu/gorm"
	"platform-sample/model"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (UserRepositoryImpl) NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

func (userRepositoryImpl UserRepositoryImpl) FindAll() ([]*model.User, error) {
	var users []*model.User
	// table.find는 할당이 안된 pointer를 넘겨도된다.
	userRepositoryImpl.db.Table("users").Find(&users)
	return users, nil
}
func (userRepositoryImpl *UserRepositoryImpl) FindById(id int) (*model.User, error) {
	var user = new(model.User)
	//userRepositoryImpl.db.Table("users").First(&user,id)
	// First는 할당안된 포인터를 넘기면 에러가 난다. nedw를 해줘야한다.
	userRepositoryImpl.db.First(&user,id)
	return user, nil

}
func (userRepositoryImpl *UserRepositoryImpl) DeleteById(id int) error {
	userRepositoryImpl.db.Delete(&model.User{},id)
	return nil
}
func (userRepositoryImpl *UserRepositoryImpl) Save(user *model.User) (*model.User, error) {
	// 여기서 예외 체크를 어떻게 할 것인가..?
	// TODO DB예외처리
	userRepositoryImpl.db.Create(user)
	return user, nil
}

func (userRepositoryImpl *UserRepositoryImpl) Update(user *model.User) (*model.User, error) {
	userRepositoryImpl.db.Save(&user)
	return user, nil
}



