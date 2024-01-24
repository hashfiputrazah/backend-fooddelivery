package service

import (
	"myapp/config"
	"myapp/model"
)

func OrderDishGetByUserIDDishID(userID, dishID string) (*model.OrderDish, error) {
	var (
		s         = config.GetDB()
		orderDish model.OrderDish
	)

	if err := s.Where("user_id = ? AND dish_id = ?", userID, dishID).Find(&orderDish).Error; err != nil {
		return nil, err
	}

	return &orderDish, nil
}

func OrderDishCreate(orderDishes []*model.OrderDish) ([]*model.OrderDish, error) {
	var s = config.GetDB()

	if err := s.Create(&orderDishes).Error; err != nil {
		return nil, err
	}

	return orderDishes, nil
}

func OrderDishGetByOrderID(orderID string) ([]*model.OrderDish, error) {
	var (
		s = config.GetDB()
		orderDishes []*model.OrderDish
	)

	if err := s.Preload("Dish").Where("order_id = ?", orderID).Find(&orderDishes).Error; err != nil {
		return nil, err
	}

	return orderDishes, nil
}
