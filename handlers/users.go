package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/maadiab/tasty-click/database"
)

func (cfg *ApiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := database.CreateUserParams{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Error in decoding data in request!!!, ", err)
		http.Error(w, "", http.StatusBadRequest)

		return
	}

	if user.Name == "" || user.Username == "" || user.Mobile == "" || user.Email == "" || user.Password == "" || user.Type == "" {
		log.Println("Error missing data!!!")
		// http.Error(w, "Error missing data!!!", http.StatusBadRequest)
		WriteError(w, 400, "Error missing data!!!")
		return
	}

	err = cfg.DBQuries.CreateUser(context.Background(), user)
	if err != nil {
		log.Println("Error in adding user to database!!!, ", err)
		http.Error(w, "Error in adding user to database", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (cfg *ApiConfig) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user_id := r.PathValue("id")
	w.Write([]byte(user_id))
}
