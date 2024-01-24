package service

import (
	"myapp/config"
	"myapp/model"

	"gorm.io/gorm"
)

func BasketGetByUserID(userID string) ([]*model.Basket, error) {
	var (
		s       = config.GetDB()
		baskets []*model.Basket
	)

	if err := s.Preload("Dish").Where("user_id = ?", userID).Find(&baskets).Error; err != nil {
		return nil, err
	}

	return baskets, nil
}

func BasketCreate(basket model.Basket) (*model.Basket, error) {
	var s = config.GetDB()

	if err := s.Omit("Dish").Create(&basket).Error; err != nil {
		return nil, err
	}

	return &basket, nil
}

func BasketDecreaseQuantity(userID, dishID string) error {
	var s = config.GetDB()

	if err := s.Model(&model.Basket{}).Where("user_id = ? AND dish_id = ?", userID, dishID).UpdateColumn("amount", gorm.Expr("amount - 1")).Error; err != nil {
		return err
	}

	return nil
}
func BasketGetByUserIDDishID(userID, dishID string) (*model.Basket, error) {
	var (
		s      = config.GetDB()
		basket model.Basket
	)

	if err := s.Where("user_id = ? AND dish_id = ?", userID, dishID).Find(&basket).Error; err != nil {
		return nil, err
	}

	return &basket, nil
}

func BasketDeleteByUserIDDishID(userID, dishID string) error {
	var s = config.GetDB()

	if err := s.Where("user_id = ? AND dish_id = ?", userID, dishID).Delete(&model.Basket{}).Error; err != nil {
		return err
	}

	return nil
}

func BasketDeleteByUserID(userID string) error {
	var s = config.GetDB()

	if err := s.Where("user_id = ?", userID).Delete(&model.Basket{}).Error; err != nil {
		return err
	}

	return nil
}

func BasketAddAmount(basketID string) error {
	var s = config.GetDB()

	if err := s.Model(&model.Basket{}).Where("id = ?", basketID).UpdateColumn("amount", gorm.Expr("amount + 1")).Error; err != nil {
		return err
	}

	return nil
}
