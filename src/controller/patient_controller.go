package controller

import (
	"encoding/json"
	"hospify/src/dao"
	"hospify/src/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePatient(w http.ResponseWriter, r *http.Request) {
	var p dao.PatientDAO
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	created, err := service.CreatePatient(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(created)
}

func GetPatients(w http.ResponseWriter, r *http.Request) {
	patients, err := service.GetAllPatients()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(patients)
}

func GetPatient(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	patient, err := service.GetPatientByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(patient)
}

func UpdatePatient(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var p dao.PatientDAO
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	p.ID = uint(id)
	updated, err := service.UpdatePatient(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

func DeletePatient(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := service.DeletePatient(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
