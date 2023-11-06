package endpoints

import (
	"gorm.io/gorm"
	"net/http"
)

type Handlers interface {
	AddOrder(w http.ResponseWriter, r *http.Request)
	CreateToken(w http.ResponseWriter, r *http.Request, userID int64, isCustomer bool) (string, error)
	DeleteOrder(w http.ResponseWriter, r *http.Request)
	GetAllOrders(w http.ResponseWriter, r *http.Request)
	ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc
	GetOrderByName(w http.ResponseWriter, r *http.Request)
}

type handlers struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) Handlers {
	return &handlers{
		db: db,
	}
}
