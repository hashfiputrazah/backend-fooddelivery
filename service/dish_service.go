package service

import (
	"myapp/config"
	"myapp/model"
)

func DishGetList(category *model.DishCategory, vegetarian bool, sorting *model.DishSorting, page int) ([]*model.Dish, error) {
	var (
		s      = config.GetDB()
		dishes []*model.Dish
		limit  = 10
		offset = (page - 1) * limit
		query  = s.Where("vegetarian = ?", vegetarian).Offset(offset).Limit(limit)
	)

	if category != nil {
		query = query.Where("category = ?", *category)
	}

	if sorting != nil {
		switch *sorting {
		case "Name Ascending":
			query = query.Order("name asc")
		case "Name Descending":
			query = query.Order("name desc")
		case "Price Ascending":
			query = query.Order("price asc")
		case "Price Descending":
			query = query.Order("price desc")
		case "Rating Ascending":
			query = query.Order("rating asc")
		case "Rating Descending":
			query = query.Order("rating desc")
		}
	}

	if err := query.Find(&dishes).Error; err != nil {
		return nil, err
	}

	return dishes, nil
}

func DishGetCount() (int, error) {
	var (
		s     = config.GetDB()
		count int64
	)

	if err := s.Model(&model.Dish{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func DishGetByID(id string) (*model.Dish, error) {
	var (
		s    = config.GetDB()
		dish model.Dish
	)

	if err := s.Where("id = ?", id).Find(&dish).Error; err != nil {
		return nil, err
	}

	return &dish, nil
}

func DishUpdateRating(id string, rating float64) error {
	var s = config.GetDB()

	if err := s.Model(&model.Dish{}).Where("id = ?", id).UpdateColumn("rating", rating).Error; err != nil {
		return err
	}

	return nil
}
