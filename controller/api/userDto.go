package api

import "platform-sample/model"

type UserDto struct {
	Name string `json:"name"`
}

func (userDto *UserDto) toModel() *model.User {
	return &model.User{Name: userDto.Name}
}
