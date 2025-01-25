package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	// "log"

	// "log"
	"net/http"

	"github.com/google/uuid"
	"github.com/maadiab/tasty-click/database"
)

func (cfg *ApiConfig) AddFood(w http.ResponseWriter, r *http.Request) {

	foods := database.CreateFoodParams{}
	if err := json.NewDecoder(r.Body).Decode(&foods); err != nil {
		// log.Println("Error in decoding food!!!",err)
		// http.Error(w,"Error in decoding food!!!",http.StatusBadRequest)
		WriteError(w, 400, "Error in decoding food!!!")
		return
	}

	if foods.Name == "" || foods.Price == 0 || foods.Picture == "" || foods.Category == "" {
		WriteError(w, 400, "Error missing data!!!")
		return
	}
	if err := cfg.DBQuries.CreateFood(context.Background(), foods); err != nil {
		WriteError(w, 500, "Error at adding food to database!!!")
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (cfg *ApiConfig) AddFoodPicture(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		WriteError(w, 405, "Error method not allowed!!!")
		return
	}

	// parse multipart form and limit size to 10 MB
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		WriteError(w, 400, "Error parsing file!!!")
		return
	}

	// retrieve file from form
	file, _, err := r.FormFile("file")
	if err != nil {
		WriteError(w, 500, "Error in reading file!!!")
		return
	}

	fileName := uuid.NewString()

	// make a file in server
	filePath := filepath.Join("./images", fileName)
	dst, err := os.Create(filePath)
	if err != nil {
		WriteError(w, 500, "Error in creating image file!!!")
		return
	}

	// copy the retived file in system file created
	_, err = io.Copy(dst, file)
	if err != nil {
		WriteError(w, 500, "Error in creating image!!!")
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("http://%s/images/%s", r.Host, fileName)))

}

func (cfg *ApiConfig) GetAllFoods(w http.ResponseWriter, r *http.Request) {

	foods, err := cfg.DBQuries.GetFoods(context.Background())
	if err != nil {
		WriteError(w, 500, "Error in getting foods!!!")
		return
	}

	data, err := json.Marshal(foods)

	if err != nil {
		WriteError(w, 500, "Error in marshalling json!!!")
		return
	}
	w.Write(data)
}
