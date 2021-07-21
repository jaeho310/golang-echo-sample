package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"platform-sample/model"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (UserRepositoryImpl) NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

func (userRepositoryImpl *UserRepositoryImpl) FindAll() ([]*model.User, error) {
	var users []*model.User
	// table.find는 할당이 안된 pointer를 넘겨도된다.
	userRepositoryImpl.db.Table("users").Find(&users)
	return users, nil
}
func (userRepositoryImpl *UserRepositoryImpl) FindById(id int) (*model.User, error) {
	var user = new(model.User)
	err := userRepositoryImpl.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil

}
func (userRepositoryImpl *UserRepositoryImpl) DeleteById(id int) error {
	err := userRepositoryImpl.db.Delete(&model.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
func (userRepositoryImpl *UserRepositoryImpl) Save(user *model.User) (*model.User, error) {
	var userCheck = new(model.User)

	// TODO 2. repository단에서 중복체크
	// 조회가 안된다는건 err가 있다는거
	err1 := userRepositoryImpl.db.Where("name = ?", user.Name).First(&userCheck).Error
	if err1 != nil {
		// 동일한 회원이 조회가 안되는경우(회원가입 가능)
		if errors.Is(err1, gorm.ErrRecordNotFound) {
			// 회원가입
			err2 := userRepositoryImpl.db.Create(user).Error
			if err2 != nil {
				return nil, err2
			}
			return user, nil
		}
	}

	// 조회가 된경우(중복된 유저이므로 회원가입 불가능)
	return nil, errors.New("중복된 유저입니다.")
}

func (userRepositoryImpl *UserRepositoryImpl) Update(user *model.User) (*model.User, error) {
	userRepositoryImpl.db.Save(&user)
	return user, nil
}

func (userRepositoryImpl *UserRepositoryImpl) DuplicatedCheck(name string) (bool, error) {
	var user = new(model.User)
	// 조회가 안된다는건 err가 있다는거
	err := userRepositoryImpl.db.Where("name = ?", name).First(&user).Error
	if err != nil {
		// 동일한 회원이 조회가 안되는경우(회원가입 가능)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		// 그 외의 에러
		return false, err
	}
	// 동일한 회원이 조회가 되는경우
	return true, errors.New("중복된 유저입니다.")
}
