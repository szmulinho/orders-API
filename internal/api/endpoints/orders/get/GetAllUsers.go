package get

import (
	"encoding/json"
	"github.com/szmulinho/orders/internal/database"
	"github.com/szmulinho/orders/internal/model"
	"net/http"
)

func GetAllOrders(w http.ResponseWriter, r *http.Request) {

	result := database.DB.Find(&model.Orders)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(model.Orders)
}
