package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/dapi/continuum-demo/internal/orders"
)

type createOrderRequest struct {
	Customer string `json:"customer"`
	Amount   int64  `json:"amount"`
}

func main() {
	svc := orders.NewService()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	mux.HandleFunc("GET /orders", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(svc.List())
	})

	mux.HandleFunc("POST /orders", func(w http.ResponseWriter, r *http.Request) {
		var input createOrderRequest
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		order, err := svc.Create(input.Customer, input.Amount)
		if err != nil {
			if errors.Is(err, orders.ErrInvalidCustomer) || errors.Is(err, orders.ErrInvalidAmount) {
				http.Error(w, err.Error(), http.StatusUnprocessableEntity)
				return
			}
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(order)
	})

	log.Println("continuum demo API listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
