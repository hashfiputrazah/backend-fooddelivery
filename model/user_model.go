package model

import "time"

type User struct {
	ID          string    `json:"id"`
	FullName    string    `json:"fullName"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	Address     string    `json:"address"`
	BirthDate   time.Time `json:"birthDate"`
	Gender      string    `json:"gender"`
	PhoneNumber string    `json:"phoneNumber"`
}

type UserRegisterParam struct {
	FullName    string `json:"fullName"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	BirthDate   string `json:"birthDate"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phoneNumber"`
}

type UserLoginParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	ID          string    `json:"id"`
	FullName    string    `json:"fullName"`
	Email       string    `json:"email"`
	Address     string    `json:"address"`
	BirthDate   time.Time `json:"birthDate"`
	Gender      string    `json:"gender"`
	PhoneNumber string    `json:"phoneNumber"`
}

type UserUpdateParam struct {
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
	BirthDate   string `json:"birthDate"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phoneNumber"`
}
