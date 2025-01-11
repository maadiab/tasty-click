package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	s := &http.Server{
		Handler: mux,
		Addr:    "localhost:8080",
	}

	s.ListenAndServe()
}
