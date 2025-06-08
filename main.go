package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"hospify/src/config"
	"hospify/src/routes"
)

func main() {
	_ = godotenv.Load()
	config.InitDB()

	r := mux.NewRouter()

	// Register all routes (modularized)
	routes.RegisterPatientRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("🚀 Server running at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
