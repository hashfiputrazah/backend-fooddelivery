package service

import (
	"myapp/config"
	"myapp/model"
)

func UserCreate(user model.User) (*model.User, error) {
	var s = config.GetDB()

	if err := s.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func UserGetByEmail(email string) (*model.User, error) {
	var (
		s    = config.GetDB()
		user model.User
	)

	if err := s.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func UserGetByID(id string) (*model.User, error) {
	var (
		s    = config.GetDB()
		user model.User
	)

	if err := s.Where("id = ?", id).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func UserUpdate(user model.User) (*model.User, error) {
	var s = config.GetDB()

	if err := s.Updates(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
