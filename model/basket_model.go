package model

type Basket struct {
	ID         string  `json:"id"`
	UserID     string  `json:"user_id"`
	DishID     string  `json:"dish_id"`
	Amount     int     `json:"amount"`
 
	Dish Dish `json:"dish" gorm:"foreignKey:DishID;references:ID"`
}

type BasketResponse struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Amount     int     `json:"amount"`
	TotalPrice float64 `json:"total_price"`
	Image      *string `json:"image"`
}
