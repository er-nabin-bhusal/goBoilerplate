package main

import (
	"fmt"
	"net/http"

	routes "github.com/nabinbhusal80/goBoilerplate/internal/routes"
)

func main() {
	router := routes.NewRouter()
	fmt.Printf("Server is running on port 8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		panic(err)
	}
}
