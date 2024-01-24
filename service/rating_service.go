package service

import (
	"myapp/config"
	"myapp/model"
)

func RatingGetByUserIDDishID(userID, dishID string) (*model.Rating, error) {
	var (
		s      = config.GetDB()
		rating model.Rating
	)

	if err := s.Where("user_id = ? AND dish_id = ?", userID, dishID).Find(&rating).Error; err != nil {
		return nil, err
	}

	return &rating, nil
}

func RatingCreate(rating model.Rating) (*model.Rating, error) {
	var s = config.GetDB()

	if err := s.Create(&rating).Error; err != nil {
		return nil, err
	}

	return &rating, nil
}

func RatingGetByDishID(dishID string) ([]*model.Rating, error) {
	var (
		s       = config.GetDB()
		ratings []*model.Rating
	)

	if err := s.Where("dish_id = ?", dishID).Find(&ratings).Error; err != nil {
		return nil, err
	}

	return ratings, nil
}
