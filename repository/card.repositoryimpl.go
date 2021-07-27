package repository

import (
	"github.com/jinzhu/gorm"
	"platform-sample/controller/dto"
	"platform-sample/model"
)

type CardRepositoryImpl struct {
	db *gorm.DB
}

func (CardRepositoryImpl) NewCardRepositoryImpl(db *gorm.DB) *CardRepositoryImpl {
	return &CardRepositoryImpl{db}
}

func (cardRepositoryImpl *CardRepositoryImpl) Save(cardDto *dto.CardDto) (*model.Card, error) {
	card := cardDto.ToModel()
	err := cardRepositoryImpl.db.Save(card).Error
	if err != nil {
		return nil, err
	}
	return card, nil
}

func (cardRepositoryImpl *CardRepositoryImpl) GetCards() (*[]model.Card, error) {
	cards := &[]model.Card{}
	err := cardRepositoryImpl.db.Table("cards").Find(&cards).Error
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func (cardRepositoryImpl *CardRepositoryImpl) DeleteById(cardId int, userId int) error {
	err := cardRepositoryImpl.db.
		Where("id = ?", cardId).
		Where("user_id = ?", userId).
		Delete(&model.Card{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (cardRepositoryImpl CardRepositoryImpl) DeleteByUserId(userId int) error {
	err := cardRepositoryImpl.db.Where("user_id = ?", userId).Delete(&model.Card{}).Error
	if err != nil {
		return err
	}
	return nil
}
