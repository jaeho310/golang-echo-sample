package service

import (
	"platform-sample/controller/dto"
	"platform-sample/model"
)

type CardService interface {
	CreateCard(*dto.CardDto) (*model.Card, error)
	DeleteCard(int, int) error
}
