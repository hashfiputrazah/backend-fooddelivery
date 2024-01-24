package model

type OrderDish struct {
	OrderID string `json:"order_id"`
	UserID  string `json:"user_id"`
	DishID  string `json:"dish_id"`
	Amount  int    `json:"amount"`

	Dish Dish `json:"dish" gorm:"foreignKey:DishID;references:ID"`
}
