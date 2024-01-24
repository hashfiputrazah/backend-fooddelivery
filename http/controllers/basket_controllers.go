package controllers

import (
	"myapp/helpers"
	"myapp/model"
	"myapp/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BasketByUser(c *gin.Context) {
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

	response := []model.BasketResponse{}
	for _, v := range baskets {
		response = append(response, model.BasketResponse{
			ID:         v.ID,
			Name:       v.Dish.Name,
			Price:      v.Dish.Price,
			Amount:     v.Amount,
			TotalPrice: v.Dish.Price * float64(v.Amount),
			Image:      v.Dish.Image,
		})
	}

	c.JSON(http.StatusOK, response)
}

func BasketAddDish(c *gin.Context) {
	dishID := c.Param("dishId")

	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}

	basket, err := service.BasketGetByUserIDDishID(userID.(string), dishID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	if basket.ID == "" {
		_, err = service.BasketCreate(model.Basket{
			ID:     helpers.UUIDGen(),
			UserID: userID.(string),
			DishID: dishID,
			Amount: 1,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Error",
				"message": err.Error(),
			})
			c.Abort()
			return
		}
	} else {
		err = service.BasketAddAmount(basket.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Error",
				"message": err.Error(),
			})
			c.Abort()
			return
		}
	}

	c.JSON(http.StatusOK, "Success")
}

func BasketDeleteDish(c *gin.Context) {
	dishID := c.Param("dishId")

	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}

	increaseParam, ok := c.GetQuery("increase")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "Parameter increase not valid",
		})
		c.Abort()
		return
	}

	increase, err := strconv.ParseBool(increaseParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	if increase {
		basket, err := service.BasketGetByUserIDDishID(userID.(string), dishID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "Error",
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		if basket.Amount > 1 {
			err = service.BasketDecreaseQuantity(userID.(string), dishID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  "Error",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
		} else {
			err = service.BasketDeleteByUserIDDishID(userID.(string), dishID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  "Error",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
		}

	} else {
		err = service.BasketDeleteByUserIDDishID(userID.(string), dishID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Error",
				"message": err.Error(),
			})
			c.Abort()
			return
		}
	}

	c.JSON(http.StatusOK, "Success")
}
