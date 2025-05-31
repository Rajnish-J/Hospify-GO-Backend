package controller

import (
	"encoding/json"
	"goBackend/Models"
	"goBackend/Service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)
 
func GetPatients(w http.ResponseWriter, r *http.Request) {
    patients, err := Service.GetAllPatients()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(patients)
}
 
func CreatePatient(w http.ResponseWriter, r *http.Request) {
    var p Models.Patient
    json.NewDecoder(r.Body).Decode(&p)
    err := Service.CreatePatient(p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}
 
func UpdatePatient(w http.ResponseWriter, r *http.Request) {
    var p Models.Patient
    json.NewDecoder(r.Body).Decode(&p)
    err := Service.UpdatePatient(p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
 
func DeletePatient(w http.ResponseWriter, r *http.Request) {
    idStr := mux.Vars(r)["id"]
    id, _ := strconv.Atoi(idStr)
    err := Service.DeletePatient(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}