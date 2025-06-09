package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"hospify/src/dto"
	"hospify/src/service"
)

type DoctorController struct {
	service *service.DoctorService
}

func NewDoctorController(service *service.DoctorService) *DoctorController {
	return &DoctorController{service: service}
}

func (c *DoctorController) CreateDoctor(w http.ResponseWriter, r *http.Request) {
	var dto dto.DoctorDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := c.service.CreateDoctor(dto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dto)
}

func (c *DoctorController) GetDoctorByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		http.Error(w, "Invalid doctor ID", http.StatusBadRequest)
		return
	}

	doctorDTO, err := c.service.GetDoctorByID(uint(id))
	if err != nil {
		http.Error(w, "Doctor not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(doctorDTO)
}

func (c *DoctorController) GetAllDoctors(w http.ResponseWriter, r *http.Request) {
	doctors, err := c.service.GetAllDoctors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(doctors)
}

func (c *DoctorController) UpdateDoctor(w http.ResponseWriter, r *http.Request) {
	var dto dto.DoctorDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := c.service.UpdateDoctor(dto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(dto)
}

func (c *DoctorController) DeleteDoctor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		http.Error(w, "Invalid doctor ID", http.StatusBadRequest)
		return
	}

	if err := c.service.DeleteDoctor(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
