package service

import (
	"platform-sample/controller/dto"
	"platform-sample/model"
	"platform-sample/repository"
)

type CardServiceImpl struct {
	repository.CardRepository
}

func (CardServiceImpl) NewCardServiceImpl(repository repository.CardRepository) *CardServiceImpl {
	return &CardServiceImpl{repository}
}

func (cardServiceImpl *CardServiceImpl) CreateCard(cardDto *dto.CardDto) (*model.Card, error) {
	return cardServiceImpl.CardRepository.Save(cardDto)
}

func (cardServiceImpl *CardServiceImpl) DeleteCard(cardId int, userId int) error {
	return cardServiceImpl.CardRepository.DeleteById(cardId, userId)
}
