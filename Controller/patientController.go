package controller

import (
	"encoding/json"
	"goBackend/Models"
	"goBackend/Service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetPatients handles GET /patients
func GetPatients(w http.ResponseWriter, r *http.Request) {
	patients, err := service.GetAllPatients()
	if err != nil {
		http.Error(w, "Failed to fetch patients: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patients)
}

// CreatePatient handles POST /patients
func CreatePatient(w http.ResponseWriter, r *http.Request) {
	var p Models.Patient
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := service.CreatePatient(p); err != nil {
		http.Error(w, "Failed to create patient: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"Patient created successfully"}`))
}

// UpdatePatient handles PUT /patients
func UpdatePatient(w http.ResponseWriter, r *http.Request) {
	var p Models.Patient
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	if p.PatientId == 0 {
		http.Error(w, "Missing patient_id in request body", http.StatusBadRequest)
		return
	}
	if err := service.UpdatePatient(p); err != nil {
		http.Error(w, "Failed to update patient: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`{"message":"Patient updated successfully"}`))
}

// DeletePatient handles DELETE /patients/{id}
func DeletePatient(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}
	if err := service.DeletePatient(id); err != nil {
		http.Error(w, "Failed to delete patient: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`{"message":"Patient deleted successfully"}`))
}
