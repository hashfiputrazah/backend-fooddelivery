package service

import (
	"myapp/config"
	"myapp/model"
)

func OrderGetByID(id string) (*model.Order, error) {
	var (
		s     = config.GetDB()
		order model.Order
	)

	if err := s.Where("id = ?", id).Find(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func OrderByUserID(userID string) ([]*model.Order, error) {
	var (
		s      = config.GetDB()
		orders []*model.Order
	)

	if err := s.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func OrderCreate(order model.Order) (*model.Order, error) {
	var s = config.GetDB()

	if err := s.Create(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func OrderUpdateStatus(id string, status model.OrderStatus) error {
	var s = config.GetDB()

	if err := s.Model(&model.Order{}).Where("id = ?", id).UpdateColumn("status", status).Error; err != nil {
		return err
	}

	return nil
}
