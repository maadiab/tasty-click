package handlers

import "net/http"

func (cfg *ApiConfig) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))
}
