package service

import (
	"doctor-booking-api/models"
	"doctor-booking-api/repository"
)

func GetDoctors() ([]models.Doctor, error) {
	return repository.GetAllDoctors()
}

func AddDoctor(d models.Doctor) (models.Doctor, error) {
	return repository.CreateDoctor(d)
}
