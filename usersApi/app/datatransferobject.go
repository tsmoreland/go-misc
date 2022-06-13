package app

import (
	"usersApi/domain"
)

type UserSummaryDto struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type UserDto struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"fullname"`
}

func NewUserDto(user domain.User) *UserDto {
	u := UserDto{ID: user.ID, Name: user.Name, FullName: user.FullName}
	return &u
}

func NewUserSummaryDto(id int64, name string) *UserSummaryDto {
	u := UserSummaryDto{ID: id, Name: name}
	return &u
}

type UserInputDto struct {
	Name     string `json:"name"`
	FullName string `json:"fullname"`
}

func (dto UserInputDto) BuildUser() *domain.User {
	user := domain.User{ID: 0, Name: dto.Name, FullName: dto.FullName}
	return &user
}

func (dto UserInputDto) UpdateUser(user domain.User) *domain.User {
	updatedUser := domain.User{ID: user.ID, Name: dto.Name, FullName: dto.FullName}
	return &updatedUser
}
