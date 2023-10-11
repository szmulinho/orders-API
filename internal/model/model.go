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

type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement"`
	Login    string `gorm:"unique" json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Exception struct {
	Message string `json:"message"`
}

var JwtKey = []byte(os.Getenv("JWT_KEY"))
