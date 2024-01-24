package controllers

import (
	"myapp/helpers"
	"myapp/model"
	"myapp/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DishList(c *gin.Context) {
	var (
		category   *model.DishCategory
		vegetarian = false
		sorting    *model.DishSorting
		page       = 1

		err error
	)

	if categoryParam, ok := c.GetQuery("category"); ok {
		temp := model.DishCategory(categoryParam)
		category = &temp
	}

	if vegetarianParam, ok := c.GetQuery("vegetarian"); ok {
		vegetarian, err = strconv.ParseBool(vegetarianParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "Error",
				"message": err.Error(),
			})
			c.Abort()
			return
		}
	}

	if sortingParam, ok := c.GetQuery("sorting"); ok {
		temp := model.DishSorting(sortingParam)
		sorting = &temp
	}

	if pageParam, ok := c.GetQuery("page"); ok {
		page, err = strconv.Atoi(pageParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "Error",
				"message": err.Error(),
			})
			c.Abort()
			return
		}
	}

	dishes, err := service.DishGetList(category, vegetarian, sorting, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	count, err := service.DishGetCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	response := model.DishPagination{
		Dishes: dishes,
		Pagination: model.Pagination{
			Size:    10,
			Count:   count,
			Current: page,
		},
	}

	c.JSON(http.StatusOK, response)
}

func DishDetail(c *gin.Context) {
	id := c.Param("id")

	dish, err := service.DishGetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, dish)
}

func DishRatingCheck(c *gin.Context) {
	id := c.Param("id")

	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}

	rating, err := service.RatingGetByUserIDDishID(userID.(string), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	if rating.ID != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": "rating has been given",
		})
		c.Abort()
		return
	}

	orderDish, err := service.OrderDishGetByUserIDDishID(userID.(string), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	if orderDish.DishID == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": "cannot give rating, never order this dish",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, "success")
}

func DishRating(c *gin.Context) {
	id := c.Param("id")

	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}

	ratingScoreParam, ok := c.GetQuery("ratingScore")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "Parameter ratingScore not valid",
		})
		c.Abort()
		return
	}

	ratingScore, err := strconv.Atoi(ratingScoreParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	if ratingScore < 1 || ratingScore > 10 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "Parameter RatingScore must between 1 - 10",
		})
		c.Abort()
		return
	}

	_, err = service.RatingCreate(model.Rating{
		ID:     helpers.UUIDGen(),
		UserID: userID.(string),
		DishID: id,
		Rating: ratingScore,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	ratings, err := service.RatingGetByDishID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	totalRating := 0.0
	for _, v := range ratings {
		totalRating += float64(v.Rating)
	}
	totalRating = totalRating / float64(len(ratings))

	err = service.DishUpdateRating(id, totalRating)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, "success")
}
