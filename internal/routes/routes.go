package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/health", healthHandler)
	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Index Handler")
	fmt.Fprintf(w, "Hello World")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Health Handler")
	fmt.Fprintf(w, "Healthy :)")
}
