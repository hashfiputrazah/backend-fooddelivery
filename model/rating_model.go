package model

type Rating struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	DishID string `json:"dish_id"`
	Rating int    `json:"rating"`
}
