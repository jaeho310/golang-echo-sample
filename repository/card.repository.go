package repository

import (
	"platform-sample/controller/dto"
	"platform-sample/model"
)

type CardRepository interface {
	Save(*dto.CardDto) (*model.Card, error)
	DeleteById(int, int) error
	DeleteByUserId(int) error
}
