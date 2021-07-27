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
	//userRepositoryImpl.db.Table("users").Find(&users)
	err := userRepositoryImpl.db.Preload("Cards").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (userRepositoryImpl *UserRepositoryImpl) FindById(id int) (*model.User, error) {
	var user = new(model.User)

	// eager loading 이며 커넥션을 두번 맺는다.(쿼리가 두번 날라간다.)
	err := userRepositoryImpl.db.Preload("Cards").First(&user, id).Error

	// join을 사용해서 map을 넘겨줄껀지, eager loading에 커넥션을 두번 맺을지는 성능을 고려하여 선택..
	//err := userRepositoryImpl.db.
	//	Table("users").
	//	Select("users.*, cards.*").
	//	Where("users.id = ?",id).
	//	Joins("join cards on users.id = cards.user_id").
	//	Scan(&user).Error

	if err != nil {
		return nil, err
	}
	return user, nil

}
func (userRepositoryImpl *UserRepositoryImpl) DeleteById(id int) error {

	userRepositoryImpl.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&model.User{}, id).Error
		if err != nil {
			return err
		}

		err = tx.Where("user_id = ?", id).Delete(&model.Card{}).Error
		if err != nil {
			return err
		}

		return nil
	})
	// 아래의 방법으로 참조를 끊어버릴수도 있다.
	//userRepositoryImpl.db.Model(&model.User{}).Association("Cards").Clear()
	// db.Model(&user).Association("Languages").Delete([]Language{languageZH, languageEN})
	// db.Model(&user).Association("Languages").Clear()

	return nil
}
func (userRepositoryImpl *UserRepositoryImpl) Save(user *model.User) (*model.User, error) {
	var userCheck = new(model.User)
	err1 := userRepositoryImpl.db.Where("name = ?", user.Name).First(&userCheck).Error
	if err1 != nil {
		if errors.Is(err1, gorm.ErrRecordNotFound) {
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

// deprecated
//func (userRepositoryImpl *UserRepositoryImpl) DuplicatedCheck(name string) (bool, error) {
//	var user = new(model.User)
//	// 조회가 안된다는건 err가 있다는거
//	err := userRepositoryImpl.db.Where("name = ?", name).First(&user).Error
//	if err != nil {
//		// 동일한 회원이 조회가 안되는경우(회원가입 가능)
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return false, nil
//		}
//		// 그 외의 에러
//		return false, err
//	}
//	// 동일한 회원이 조회가 되는경우
//	return true, errors.New("중복된 유저입니다.")
//}
