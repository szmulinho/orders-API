package server

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/szmulinho/orders/internal/api/endpoints/orders/add"
	"github.com/szmulinho/orders/internal/api/endpoints/orders/delete"
	"github.com/szmulinho/orders/internal/api/endpoints/orders/get"
	"github.com/szmulinho/orders/internal/api/jwt"
	"log"
	"net/http"
)

func Run() {

	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/order", add.AddOrder).Methods("POST")
	router.HandleFunc("/orders", get.GetAllOrders).Methods("GET")
	router.HandleFunc("/delete_order/{id}", delete.DeleteOrder).Methods("DELETE")
	router.HandleFunc("/authenticate", func(w http.ResponseWriter, r *http.Request) {
		userID := uint(1)
		isDoctor := true
		token, err := jwt.CreateToken(w, r, int64(userID), isDoctor)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(token))
	}).Methods("POST")
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "Content-Type"}),
		handlers.ExposedHeaders([]string{}),
		handlers.AllowCredentials(),
		handlers.MaxAge(86400),
	)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8084"), cors(router)))

}

func Server() {
	log.Println("server with port 8081 is starting")
	Run()
}
