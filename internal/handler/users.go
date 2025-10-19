package handler

import "net/http"

func (cfg *ApiConfig) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getusers"))
}
