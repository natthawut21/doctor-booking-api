package repository

import (
	"doctor-booking-api/config"
	"doctor-booking-api/models"
)

func GetAllDoctors() ([]models.Doctor, error) {
	var doctors []models.Doctor
	result := config.DB.Find(&doctors)
	return doctors, result.Error
}

func CreateDoctor(doctor models.Doctor) (models.Doctor, error) {
	result := config.DB.Create(&doctor)
	return doctor, result.Error
}
