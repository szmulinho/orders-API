package endpoints

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/szmulinho/orders/internal/model"
	"net/http"
)

func (h *handlers) GetOrderByName(w http.ResponseWriter, r *http.Request) {
	patient := mux.Vars(r)["name"]

	if err := h.db.Where("name = ?", patient).Find(&model.Orders).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(model.Orders)
}
