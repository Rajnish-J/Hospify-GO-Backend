package main

import (
	"fmt"
	"goBackend/config"
	"goBackend/controller"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
    config.InitDB()

    r := mux.NewRouter()

    r.HandleFunc("/patients", controller.GetPatients).Methods("GET")
    r.HandleFunc("/patients", controller.CreatePatient).Methods("POST")
    r.HandleFunc("/patients", controller.UpdatePatient).Methods("PUT")
    r.HandleFunc("/patients/{id}", controller.DeletePatient).Methods("DELETE")

    fmt.Println("🚀 Server running at http://localhost:8080")

    // CORS middleware (optional, but important if connecting with frontend)
    headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
    origins := handlers.AllowedOrigins([]string{"*"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

    err := http.ListenAndServe(":8080", handlers.CORS(origins, headers, methods)(r))
    if err != nil {
        log.Fatal("Failed to start server: ", err)
    }
}
