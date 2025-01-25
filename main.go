package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/maadiab/tasty-click/database"
	"github.com/maadiab/tasty-click/handlers"
	"log"
	"net/http"
	"os"
	// "github.com/pressly/goose/v3/database"
)

func main() {

	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Println("Error in connecting database!!!, ", err)
		return
	}
	cfg := &handlers.ApiConfig{
		DBQuries: *database.New(db),
	}

	mux := http.NewServeMux()
	mux.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
	mux.HandleFunc("/", cfg.Home)
	mux.HandleFunc("POST /api/users", cfg.CreateUser)
	mux.HandleFunc("POST /api/users/{id}", cfg.UpdateUser)
	mux.HandleFunc("GET /api/foods", cfg.GetAllFoods)
	mux.HandleFunc("POST /api/foods", cfg.AddFood)
	mux.HandleFunc("POST /api/foodphoto", cfg.AddFoodPicture)
	mux.HandleFunc("POST /api/orders", cfg.CreateOrder)
	s := &http.Server{
		Handler: mux,
		Addr:    "localhost:5000",
	}
	log.Println("Server is listining on: ", s.Addr)
	s.ListenAndServe()

}
