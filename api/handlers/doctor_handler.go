package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/alifhaider/carehub-backend/internal/database"
	"github.com/alifhaider/carehub-backend/internal/models"
)

func CreateDoctor(w http.ResponseWriter, r *http.Request) {
    var doctor models.Doctor
    if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    database.DB.Create(&doctor)
    json.NewEncoder(w).Encode(doctor)
}

func GetDoctors(w http.ResponseWriter, r *http.Request) {
    var doctors []models.Doctor
    database.DB.Find(&doctors)
    json.NewEncoder(w).Encode(doctors)
}

func GetDoctorByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var doctor models.Doctor
    database.DB.First(&doctor, id)
    if doctor.ID == 0 {
        http.Error(w, "Doctor not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(doctor)
}

func UpdateDoctor(w http.ResponseWriter, r *http.Request) {
    var doctor models.Doctor
    if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    database.DB.Save(&doctor)
    json.NewEncoder(w).Encode(doctor)
}

func DeleteDoctor(w http.ResponseWriter, r *http.Request) {
    var doctor models.Doctor
    if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    database.DB.Delete(&doctor)
    w.WriteHeader(http.StatusNoContent)
}

func GetDoctorsBySpecialization(w http.ResponseWriter, r *http.Request) {
    specialization := r.URL.Query().Get("specialization")
    var doctors []models.Doctor
    database.DB.Where("specialization = ?", specialization).Find(&doctors)
    json.NewEncoder(w).Encode(doctors)
}

func GetDoctorsByLocation(w http.ResponseWriter, r *http.Request) {
    location := r.URL.Query().Get("location")
    var doctors []models.Doctor
    database.DB.Where("location = ?", location).Find(&doctors)
    json.NewEncoder(w).Encode(doctors)
}

func GetDoctorsByAvailability(w http.ResponseWriter, r *http.Request) {
	availability := r.URL.Query().Get("availability")
	var doctors []models.Doctor
	database.DB.Where("availability = ?", availability).Find(&doctors)
	json.NewEncoder(w).Encode(doctors)
}

func GetDoctorsByRating(w http.ResponseWriter, r *http.Request) {
	rating, err := strconv.ParseFloat(r.URL.Query().Get("rating"), 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var doctors []models.Doctor
	database.DB.Where("rating >= ?", rating).Find(&doctors)
	json.NewEncoder(w).Encode(doctors)
}


