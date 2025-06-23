package service

import (
	"doctor-booking-api/models"
	"doctor-booking-api/repository"
	"errors"
	"strings"
	"time"
)

// ตรวจสอบและบันทึกตารางใหม่
func SaveSchedule(request models.DoctorSchedule) (models.DoctorSchedule, error) {
	// ตรวจสอบความถูกต้องของ dayOfWeek
	if !isValidDayOfWeek(request.DayOfWeek) {
		return models.DoctorSchedule{}, errors.New("Invalid dayOfWeek: must be MONDAY to SUNDAY")
	}

	// ตรวจสอบเวลาเริ่มต้องน้อยกว่าจบ
	if !isValidTimeRange(request.StartTime, request.EndTime) {
		return models.DoctorSchedule{}, errors.New("Invalid time range: endTime must be after startTime")
	}

	return repository.CreateDoctorSchedule(request)
}

// อัปเดตตารางทำงาน
func UpdateScheduleByID(scheduleID uint, updated models.DoctorSchedule) (models.DoctorSchedule, error) {
	if !isValidDayOfWeek(updated.DayOfWeek) {
		return models.DoctorSchedule{}, errors.New("Invalid dayOfWeek")
	}
	if !isValidTimeRange(updated.StartTime, updated.EndTime) {
		return models.DoctorSchedule{}, errors.New("Invalid time range")
	}
	return repository.UpdateDoctorScheduleByID(scheduleID, updated)
}

// ดึงตารางทั้งหมดของหมอ
func GetSchedulesByDoctorID(doctorID uint) ([]models.DoctorSchedule, error) {
	return repository.GetSchedulesByDoctorID(doctorID)
}

// ลบตารางของหมอแบบเฉพาะเจาะจง
func DeleteScheduleByDoctorAndID(doctorID uint, scheduleID uint) error {
	return repository.DeleteScheduleByDoctorAndID(doctorID, scheduleID)
}

// ฟังก์ชันช่วยตรวจสอบ dayOfWeek (MONDAY–SUNDAY)
func isValidDayOfWeek(day string) bool {
	days := []string{
		"MONDAY", "TUESDAY", "WEDNESDAY",
		"THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY",
	}
	day = strings.ToUpper(day)
	for _, d := range days {
		if d == day {
			return true
		}
	}
	return false
}

// ตรวจสอบว่า endTime ต้องมากกว่า startTime
func isValidTimeRange(start string, end string) bool {
	layout := "15:04"
	startTime, err1 := time.Parse(layout, start)
	endTime, err2 := time.Parse(layout, end)
	if err1 != nil || err2 != nil {
		return false
	}
	return endTime.After(startTime)
}
