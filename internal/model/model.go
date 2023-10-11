package model

import (
	"os"
)

type Order struct {
	ID      int64  `json:"order-id" gorm:"primaryKey;autoIncrement"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Items   string `json:"items"`
	Price   string `json:"price"`
}

var Orders []Order

var NewOrder Order

type Port struct {
	Port string
}

type jwtUser struct {
	Jwtuser     string "jwt-user"
	Jwtpassword string "jwtpassword"
}

var Juser jwtUser

type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement"`
	Login    string `gorm:"unique" json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var Users []User

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

type Response struct {
	Data string `json:"data"`
}

var JwtKey = []byte(os.Getenv("JWT_KEY"))

type LoginResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
