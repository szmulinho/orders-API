package endpoints

import (
	"encoding/json"
	"github.com/szmulinho/orders/internal/model"
	"net/http"
)

func (h *handlers) GetAllOrders(w http.ResponseWriter, r *http.Request) {

	result := h.db.Find(&model.Orders)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(model.Orders)
}
