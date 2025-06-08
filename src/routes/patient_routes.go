package routes

import (
	"github.com/gorilla/mux"
	"hospify/src/controller"
)

func RegisterPatientRoutes(router *mux.Router) {
	router.HandleFunc("/patients", controller.CreatePatient).Methods("POST")
	router.HandleFunc("/patients", controller.GetPatients).Methods("GET")
	router.HandleFunc("/patients/{id}", controller.GetPatient).Methods("GET")
	router.HandleFunc("/patients/{id}", controller.UpdatePatient).Methods("PUT")
	router.HandleFunc("/patients/{id}", controller.DeletePatient).Methods("DELETE")
}
