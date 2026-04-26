package main

import (
	"fmt"
	"net/http"

	"github.com/marifyahya/todo-app/backend/database"
)

func main() {
	// Initialize Database
	database.InitDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Todos Backend!")
	})

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
