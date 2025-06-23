// service/appointment_service.go
package service

import (
	"doctor-booking-api/config"
	"doctor-booking-api/models"
	"errors"
	"time"
)

type BookRequest struct {
	DoctorID uint   `json:"doctor_id"`
	SlotID   uint   `json:"slot_id"`
	Username string `json:"username"` // หา Patient จาก Username
}

func BookAppointment(req BookRequest) (uint, error) {
	// 1. หา Slot
	var slot models.AppointmentSlot
	if err := config.DB.First(&slot, req.SlotID).Error; err != nil {
		return 0, errors.New("slot not found")
	}

	// 2. เช็คว่า slot ถูกจองแล้วหรือยัง
	if slot.Booked {
		return 0, errors.New("slot already booked")
	}

	// 3. หา Doctor
	var doctor models.Doctor
	if err := config.DB.First(&doctor, req.DoctorID).Error; err != nil {
		return 0, errors.New("doctor not found")
	}

	// 4. หา Patient จาก Username
	var patient models.User
	if err := config.DB.Where("username = ?", req.Username).First(&patient).Error; err != nil {
		return 0, errors.New("patient not found")
	}

	// 5. จอง slot และสร้าง appointment
	slot.Booked = true
	config.DB.Save(&slot)

	appointment := models.Appointment{
		DoctorID:  doctor.ID,
		PatientID: patient.ID,
		SlotID:    slot.ID,
		CreatedAt: time.Now(),
	}

	if err := config.DB.Create(&appointment).Error; err != nil {
		return 0, errors.New("failed to create appointment")
	}

	return appointment.ID, nil
}

func ListAppointments() ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := config.DB.Preload("Doctor").Preload("Patient").Preload("Slot").Find(&appointments).Error
	return appointments, err
}
