package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/szmulinho/orders/internal/model"
	"log"
	"net/http"
	"net/smtp"
)

type errResponse struct {
	Error string `json:"error"`
}

func (h *handlers) AddOrder(w http.ResponseWriter, r *http.Request) {

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

	result := h.db.Create(&newOrder)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	for _, singleOrder := range model.Orders {
		fmt.Println(singleOrder)
		if singleOrder.ID == model.NewOrder.ID {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(errResponse{Error: fmt.Sprintf("Order %model.already exist", model.NewOrder.ID)})
			return
		}
	}

	fmt.Printf("created new prescription %+v\n", model.Orders)
	log.Printf("%+v", model.NewOrder)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(model.NewOrder)

	emailSubject := "Order confirmation"
	emailBody := fmt.Sprintf("Thank you for your purchase. Your order ID: %s", newOrder.ID)
	err = sendEmail(newOrder.Email, emailSubject, emailBody)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		// Handle the email sending error, maybe return an error response to the client
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)
}

func sendEmail(to, subject, body string) error {
	from := "szmul-med@outlook.com"
	password := "L96a1prosniper"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail("smtp.example.com:587", smtp.PlainAuth("", from, password, "smtp.office365.com"), from, []string{to}, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}
