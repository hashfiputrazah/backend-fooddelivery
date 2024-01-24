package model

import "time"

type Order struct {
	ID           string      `json:"id"`
	UserID       string      `json:"user_id"`
	DeliveryTime time.Time   `json:"deliveryTime"`
	OrderTime    time.Time   `json:"orderTime"`
	Status       OrderStatus `json:"status"`
	Price        float64     `json:"price"`
	Address      string      `json:"address"`
}

type OrderStatus string

const (
	InProcess OrderStatus = "In Process"
	Delivered OrderStatus = "Delivered"
)

type OrderDetailResponse struct {
	ID           string      `json:"id"`
	DeliveryTime time.Time   `json:"deliveryTime"`
	OrderTime    time.Time   `json:"orderTime"`
	Status       OrderStatus `json:"status"`
	Price        float64     `json:"price"`
	Address      string      `json:"address"`

	Dishes []*BasketResponse `json:"dishes"`
}

type OrderListResponse struct {
	ID           string      `json:"id"`
	DeliveryTime time.Time   `json:"deliveryTime"`
	OrderTime    time.Time   `json:"orderTime"`
	Status       OrderStatus `json:"status"`
	Price        float64     `json:"price"`
}

type OrderCreateParam struct {
	DeliveryTime time.Time `json:"deliveryTime"`
	Address      string    `json:"address"`
}
