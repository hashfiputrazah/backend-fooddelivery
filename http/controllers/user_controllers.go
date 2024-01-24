package controllers

import (
	"fmt"
	"myapp/helpers"
	"myapp/model"
	"myapp/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	param := model.UserRegisterParam{}
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	birthDate, err := time.Parse("2006-01-02", param.BirthDate)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "Parameter birth date is not valid",
		})
		c.Abort()
		return
	}

	if !helpers.IsValidPhoneNumber(param.PhoneNumber) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "Parameter phone number is not valid",
		})
		c.Abort()
		return
	}

	user, err := service.UserCreate(model.User{
		ID:          helpers.UUIDGen(),
		FullName:    param.FullName,
		Password:    helpers.HashPassword(param.Password),
		Email:       param.Email,
		Address:     param.Address,
		BirthDate:   birthDate,
		Gender:      param.Gender,
		PhoneNumber: param.PhoneNumber,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	jwtToken, err := helpers.GenerateJWTToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	response := model.TokenResponse{
		Token: jwtToken,
	}

	c.JSON(http.StatusOK, response)
}

func UserLogin(c *gin.Context) {
	param := model.UserLoginParam{}
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	user, err := service.UserGetByEmail(param.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	if !helpers.CheckPasswordHash(param.Password, user.Password) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": "Email or password not match",
		})
		c.Abort()
		return
	}

	jwtToken, err := helpers.GenerateJWTToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	response := model.TokenResponse{
		Token: jwtToken,
	}

	c.JSON(http.StatusOK, response)
}

func UserLogout(c *gin.Context) {
	c.JSON(http.StatusOK, "Success")
}

func UserDetail(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}

	user, err := service.UserGetByID(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	response := model.UserResponse{
		ID:          user.ID,
		FullName:    user.FullName,
		Email:       user.Email,
		Address:     user.Address,
		BirthDate:   user.BirthDate,
		Gender:      user.Gender,
		PhoneNumber: user.PhoneNumber,
	}

	c.JSON(http.StatusOK, response)
}

func UserUpdate(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}

	param := model.UserUpdateParam{}
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	if !helpers.IsValidPhoneNumber(param.PhoneNumber) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "Parameter phone number is not valid",
		})
		c.Abort()
		return
	}

	birthDate, err := time.Parse("2006-01-02", param.BirthDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "Parameter birth date is not valid",
		})
		c.Abort()
		return
	}

	user, err := service.UserGetByID(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	user.FullName = param.FullName
	user.Address = param.Address
	user.BirthDate = birthDate
	user.Gender = param.Gender
	user.PhoneNumber = param.PhoneNumber

	_, err = service.UserUpdate(*user)
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
