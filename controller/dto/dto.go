package dto

import "platform-sample/model"

type UserDto struct {
	Name string `json:"name"`
}

func (UserDto *UserDto) ToModel() *model.User {
	return &model.User{Name: UserDto.Name}
}

type CardDto struct {
	Name   string `json:"name"`
	Limit  int    `json:"limit"`
	UserId uint   `json:"userId"`
}

func (cardDto *CardDto) ToModel() *model.Card {
	return &model.Card{Name: cardDto.Name, Limit: cardDto.Limit, UserId: cardDto.UserId}
}
