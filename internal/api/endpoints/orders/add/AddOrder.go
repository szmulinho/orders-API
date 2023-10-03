package add

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/szmulinho/orders/internal/database"
	"github.com/szmulinho/orders/internal/model"
	"log"
	"net/http"
)

type errResponse struct {
	Error string `json:"error"`
}

func AddOrder(w http.ResponseWriter, r *http.Request) {

	var newOrder model.Order

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
	err = json.Unmarshal(buf.Bytes(), &newOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
		log.Printf("Invalid body")
	}

	result := database.DB.Create(&newOrder)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	for _, singleOrder := range model.Orders {
		fmt.Println(singleOrder)
		if singleOrder.OrderID == model.NewOrder.OrderID {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(errResponse{Error: fmt.Sprintf("Order %model.already exist", model.NewOrder.OrderID)})
			return
		}
	}

	fmt.Printf("created new prescription %+v\n", model.Orders)
	log.Printf("%+v", model.NewOrder)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(model.NewOrder)
}
