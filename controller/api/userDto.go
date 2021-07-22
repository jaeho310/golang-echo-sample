package api

import "platform-sample/model"

type UserDto struct {
	Name string `json:"name"`
}

func (UserDto *UserDto) toModel() *model.User{
	return &model.User{Name: UserDto.Name}
}