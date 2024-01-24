package model

type Dish struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description *string      `json:"description"`
	Price       float64      `json:"price"`
	Image       *string      `json:"image"`
	Vegetarian  bool         `json:"vegetarian"`
	Rating      float64      `json:"rating"`
	Category    DishCategory `json:"category"`
}

type DishCategory string

const (
	Wok     DishCategory = "Wok"
	Pizza   DishCategory = "Pizza"
	Soup    DishCategory = "Soup"
	Dessert DishCategory = "Dessert"
	Drink   DishCategory = "Drink"
)

type DishSorting string

const (
	NameAsc    DishSorting = "Name Ascending"
	NameDesc   DishSorting = "Name Descending"
	PriceAsc   DishSorting = "Price Ascending"
	PriceDesc  DishSorting = "Price Descending"
	RatingAsc  DishSorting = "Rating Ascending"
	RatingDesc DishSorting = "Rating Descending"
)

type DishPagination struct {
	Dishes     []*Dish    `json:"dishes"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Size    int `json:"size"`
	Count   int `json:"count"`
	Current int `json:"current"`
}
