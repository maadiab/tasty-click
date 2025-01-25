package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/maadiab/tasty-click/database"
)

func (cfg *ApiConfig) CreateOrder(w http.ResponseWriter, r *http.Request) {
	order := database.CreateOrderParams{}

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {

		WriteError(w, 400, "Error in decoding json!!!")
		return
	}

	err := cfg.DBQuries.CreateOrder(context.Background(), order)
	if err != nil {
		WriteError(w, 500, "Error creating order")
		return
	}

	w.WriteHeader(http.StatusOK)
}
