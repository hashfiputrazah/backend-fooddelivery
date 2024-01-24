package controllers

import (
	"myapp/helpers"
	"myapp/model"
	"myapp/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func OrderDetail(c *gin.Context) {
	id := c.Param("id")

	order, err := service.OrderGetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	orderDishes, err := service.OrderDishGetByOrderID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	dishes := []*model.BasketResponse{}
	for _, v := range orderDishes {
		dishes = append(dishes, &model.BasketResponse{
			ID:         v.DishID,
			Name:       v.Dish.Name,
			Price:      v.Dish.Price,
			Amount:     v.Amount,
			TotalPrice: v.Dish.Price * float64(v.Amount),
			Image:      v.Dish.Image,
		})
	}

	response := model.OrderDetailResponse{
		ID:           order.ID,
		DeliveryTime: order.DeliveryTime,
		OrderTime:    order.OrderTime,
		Status:       order.Status,
		Price:        order.Price,
		Address:      order.Address,
		Dishes:       dishes,
	}

	c.JSON(http.StatusOK, response)
}

func OrderByUser(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}

	orders, err := service.OrderByUserID(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	response := []*model.OrderListResponse{}
	for _, v := range orders {
		response = append(response, &model.OrderListResponse{
			ID:           v.ID,
			DeliveryTime: v.DeliveryTime,
			OrderTime:    v.OrderTime,
			Status:       v.Status,
			Price:        v.Price,
		})
	}

	c.JSON(http.StatusOK, response)
}

func OrderCreate(c *gin.Context) {
	orderCreateParam := model.OrderCreateParam{}
	if err := c.ShouldBindJSON(&orderCreateParam); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	currentTime := time.Now()
	if orderCreateParam.DeliveryTime.Before(currentTime) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": "Parameter delivery time must be after order time",
		})
		c.Abort()
		return
	}

	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}

	baskets, err := service.BasketGetByUserID(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	if len(baskets) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": "Cannot create order, basket is empty",
		})
		c.Abort()
		return
	}

	price := 0.0
	for _, v := range baskets {
		price += v.Dish.Price * float64(v.Amount)
	}

	order, err := service.OrderCreate(model.Order{
		ID:           helpers.UUIDGen(),
		UserID:       userID.(string),
		DeliveryTime: orderCreateParam.DeliveryTime,
		OrderTime:    time.Now(),
		Status:       model.InProcess,
		Price:        price,
		Address:      orderCreateParam.Address,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	orderDishes := []*model.OrderDish{}
	for _, v := range baskets {
		orderDishes = append(orderDishes, &model.OrderDish{
			UserID:  userID.(string),
			OrderID: order.ID,
			DishID:  v.DishID,
			Amount:  v.Amount,
		})
	}

	_, err = service.OrderDishCreate(orderDishes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	err = service.BasketDeleteByUserID(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, "Success")
}

func OrderConfirm(c *gin.Context) {
	id := c.Param("id")

	order, err := service.OrderGetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	if order.Status != model.InProcess {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": "cannot confirm order, please check order status",
		})
		c.Abort()
		return
	}

	err = service.OrderUpdateStatus(id, model.Delivered)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, "Success")
}
