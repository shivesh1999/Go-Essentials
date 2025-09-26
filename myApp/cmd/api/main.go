package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shivesh1999/myapp/internal/handler"
	"github.com/shivesh1999/myapp/internal/service"
)

func main() {
	// Initialize dependencies
	userService := service.UserService{}
	userHandler := handler.UserHandler{Service: userService}

	// Routes
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})
	http.HandleFunc("/users", userHandler.GetUser)

	// Start server
	fmt.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
