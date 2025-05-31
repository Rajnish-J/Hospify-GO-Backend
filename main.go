package main
 
import (
    "fmt"
    "goBackend/Config"
    "goBackend/Controller"
    "net/http"
 
    "github.com/gorilla/mux"
)
 
func main() {
    config.InitDB()
 
    r := mux.NewRouter()
    r.HandleFunc("/patients", controller.GetPatients).Methods("GET")
    r.HandleFunc("/patients", controller.CreatePatient).Methods("POST")
    r.HandleFunc("/patients", controller.UpdatePatient).Methods("PUT")
    r.HandleFunc("/patients/{id}", controller.DeletePatient).Methods("DELETE")
 
    fmt.Println("Server is running at http://localhost:8080")
    http.ListenAndServe(":8080", r)
}