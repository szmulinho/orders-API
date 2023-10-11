package endpoints

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/szmulinho/orders/internal/model"
	"net/http"
	"strconv"
)

func (h *handlers) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderIDStr := mux.Vars(r)["id"]
	orderID, err := strconv.ParseInt(orderIDStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, singleOrder := range model.Orders {
		if singleOrder.ID == orderID {
			model.Orders = append(model.Orders[:i], model.Orders[i+1:]...)
			fmt.Fprintf(w, "Order with ID %v has been deleted successfully", orderID)
		}
	}
}
