package repository

import (
	"doctor-booking-api/config"
	"doctor-booking-api/models"
	"errors"
)

func CreateDoctorSchedule(schedule models.DoctorSchedule) (models.DoctorSchedule, error) {
	result := config.DB.Create(&schedule)
	return schedule, result.Error
}

func UpdateDoctorScheduleByID(id uint, data models.DoctorSchedule) (models.DoctorSchedule, error) {
	var schedule models.DoctorSchedule
	
	if err := config.DB.First(&schedule, id).Error; err != nil {
		return schedule, errors.New("Schedule not found")
	}
	
	schedule.DayOfWeek = data.DayOfWeek
	schedule.StartTime = data.StartTime
	schedule.EndTime = data.EndTime

	
	if err := config.DB.Save(&schedule).Error; err != nil {
		return schedule, err
	}

	return schedule, nil
}


func GetSchedulesByDoctorID(doctorID uint) ([]models.DoctorSchedule, error) {
	var schedules []models.DoctorSchedule
	result := config.DB.Where("doctor_id = ?", doctorID).Find(&schedules)
	return schedules, result.Error
}


func DeleteScheduleByDoctorAndID(doctorID uint, scheduleID uint) error {
	var schedule models.DoctorSchedule
	result := config.DB.Where("id = ? AND doctor_id = ?", scheduleID, doctorID).First(&schedule)
	if result.Error != nil {
		return errors.New("Schedule not found")
	}
	return config.DB.Delete(&schedule).Error
}
