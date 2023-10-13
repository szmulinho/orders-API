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

type Exception struct {
	Message string `json:"message"`
}

var JwtKey = []byte(os.Getenv("JWT_KEY"))
