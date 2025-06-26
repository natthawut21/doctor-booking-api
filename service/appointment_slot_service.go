package service

import (
	"doctor-booking-api/config"
	"doctor-booking-api/models"
	"errors"
	"fmt"
	"strings"
	"time"
	"sort"
)


type SlotResponse struct {
	ID        int64      `json:"id"`
	DoctorID  int64      `json:"doctor_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    string    `json:"status"`
	ChangedAt  *time.Time `json:"changed_at,omitempty"` // เวลาที่สถานะล่าสุดถูกเปลี่ยน (nullable)
}
type SlotWithHistoryResponse struct {
	ID        int64      `json:"id"`
	DoctorID  int64      `json:"doctor_id"`
	StartTime time.Time  `json:"start_time"`
	EndTime   time.Time  `json:"end_time"`
	Status    string     `json:"status"`
	History   []models.SlotStatusHistory `json:"history"`
}
func toSlotResponse(slot models.AppointmentSlot) SlotResponse {

	var changedAt *time.Time = nil

	// ถ้าสถานะคือ PENDING หรือ CONFIRMED
	if slot.Status == "PENDING" || slot.Status == "CONFIRMED" || slot.Status == "CANCELED" {
		var latestHistory models.SlotStatusHistory
		err := config.DB.Where("slot_id = ? AND new_status = ?", slot.ID, slot.Status).
			Order("changed_at desc").
			First(&latestHistory).Error
		if err == nil {
			changedAt = &latestHistory.ChangedAt
		}
	}
	return SlotResponse{
		ID:        slot.ID,
		DoctorID:  slot.DoctorID,
		StartTime: slot.StartTime.In(time.FixedZone("Asia/Bangkok", 7*60*60)),
		EndTime:   slot.EndTime.In(time.FixedZone("Asia/Bangkok", 7*60*60)),
		Status:    slot.Status,
		ChangedAt: changedAt,
	}
}

var location, _ = time.LoadLocation("Asia/Bangkok")

func GenerateSlots(doctorID int64, date string) ([]SlotResponse, error) {
	localDate, err := time.ParseInLocation("2006-01-02", date, location)
	if err != nil {
		return nil, errors.New("Invalid date format, expected yyyy-MM-dd")
	}

	var doctor models.Doctor
	if err := config.DB.First(&doctor, doctorID).Error; err != nil {
		return nil, errors.New("Doctor not found")
	}

	dayOfWeek := strings.ToUpper(localDate.Weekday().String())

	var schedules []models.DoctorSchedule
	if err := config.DB.Where("doctor_id = ? AND day_of_week = ?", doctorID, dayOfWeek).Find(&schedules).Error; err != nil {
		return nil, fmt.Errorf("No schedule found for doctor on %s", dayOfWeek)
	}
	if len(schedules) == 0 {
		return nil, fmt.Errorf("Doctor has no schedule on %s", dayOfWeek)
	}

	var generated []SlotResponse

	for _, schedule := range schedules {
		startTime, _ := time.Parse("15:04:05", schedule.StartTime)
		endTime, _ := time.Parse("15:04:05", schedule.EndTime)

		start := time.Date(localDate.Year(), localDate.Month(), localDate.Day(), startTime.Hour(), startTime.Minute(), 0, 0, location)
		end := time.Date(localDate.Year(), localDate.Month(), localDate.Day(), endTime.Hour(), endTime.Minute(), 0, 0, location)

		for current := start; current.Add(15*time.Minute).Before(end) || current.Add(15*time.Minute).Equal(end); current = current.Add(20 * time.Minute) {
			slotStart := current
			slotEnd := current.Add(15 * time.Minute)

			var existing models.AppointmentSlot
			err := config.DB.
				Where("doctor_id = ? AND start_time = ? AND end_time = ?", doctorID, slotStart, slotEnd).
				First(&existing).Error

			if err == nil {
				generated = append(generated, toSlotResponse(existing))
			} else {
				newSlot := models.AppointmentSlot{
					DoctorID:  doctorID,
					StartTime: slotStart,
					EndTime:   slotEnd,
					Status:    "AVAILABLE",
				}
				config.DB.Create(&newSlot)
				generated = append(generated, toSlotResponse(newSlot))
			}
		}
	}

	// เรียงตามเวลา
	sort.Slice(generated, func(i, j int) bool {
		return generated[i].StartTime.Before(generated[j].StartTime)
	})

	return generated, nil
}

func ShowAllSlots(doctorID int64, dateStr string) ([]SlotResponse, error) {
	localDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, errors.New("Invalid date format, expected yyyy-MM-dd")
	}

	startOfDay := time.Date(localDate.Year(), localDate.Month(), localDate.Day(), 0, 0, 0, 0, location)
	endOfDay := startOfDay.Add(24 * time.Hour)

	var slots []models.AppointmentSlot
	err = config.DB.Where("doctor_id = ? AND start_time >= ? AND start_time < ?", doctorID, startOfDay, endOfDay).
		Order("start_time asc").Find(&slots).Error
	if err != nil {
		return nil, err
	}


	if len(slots) == 0 {
		return GenerateSlots(doctorID, dateStr)
	}

	
	var responses []SlotResponse
	for _, s := range slots {
		responses = append(responses, toSlotResponse(s))
	}

	
	sort.Slice(responses, func(i, j int) bool {
		return responses[i].StartTime.Before(responses[j].StartTime)
	})

	return responses, nil
}


func AvailableSlots(doctorID int64, dateStr string) ([]SlotResponse, error) {
	localDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, errors.New("Invalid date format, expected yyyy-MM-dd")
	}

	startOfDay := time.Date(localDate.Year(), localDate.Month(), localDate.Day(), 0, 0, 0, 0, location)
	endOfDay := startOfDay.Add(24 * time.Hour)

	// ตรวจสอบว่ามี slot แล้วหรือยัง
	var existing []models.AppointmentSlot
	err = config.DB.Where("doctor_id = ? AND start_time >= ? AND start_time < ?", doctorID, startOfDay, endOfDay).
		Find(&existing).Error
	if err != nil {
		return nil, err
	}
	if len(existing) == 0 {
		_, err := GenerateSlots(doctorID, dateStr)
		if err != nil {
			return nil, err
		}
	}

	
	var available []models.AppointmentSlot
	err = config.DB.Where("doctor_id = ? AND start_time >= ? AND start_time < ? AND status = ?", doctorID, startOfDay, endOfDay, "AVAILABLE").
		Order("start_time asc").Find(&available).Error
	if err != nil {
		return nil, err
	}

	
	var responses []SlotResponse
	for _, s := range available {
		responses = append(responses, toSlotResponse(s))
	}

	return responses, nil
}
func BookedSlots(doctorID int64, dateStr string) ([]SlotResponse, error) {
	localDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, errors.New("Invalid date format, expected yyyy-MM-dd")
	}

	startOfDay := time.Date(localDate.Year(), localDate.Month(), localDate.Day(), 0, 0, 0, 0, location)
	endOfDay := startOfDay.Add(24 * time.Hour)

	// ตรวจสอบว่ามี slot แล้วหรือยัง
	var existing []models.AppointmentSlot
	err = config.DB.Where("doctor_id = ? AND start_time >= ? AND start_time < ?", doctorID, startOfDay, endOfDay).
		Find(&existing).Error
	if err != nil {
		return nil, err
	}
	if len(existing) == 0 {
		_, err := GenerateSlots(doctorID, dateStr)
		if err != nil {
			return nil, err
		}
	}

	
	var available []models.AppointmentSlot
	err = config.DB.Where("doctor_id = ? AND start_time >= ? AND start_time < ? AND status = ?", doctorID, startOfDay, endOfDay, "CONFIRMED").
		Order("start_time asc").Find(&available).Error
	if err != nil {
		return nil, err
	}

	
	var responses []SlotResponse
	for _, s := range available {
		responses = append(responses, toSlotResponse(s))
	}

	return responses, nil
}
func UpdateSlotStatus(slotID int64, newStatus string, changedBy string) error {
	var slot models.AppointmentSlot
	if err := config.DB.First(&slot, slotID).Error; err != nil {
		return errors.New("slot not found")
	}

	validStatuses := map[string]bool{
		"AVAILABLE": true,
		"PENDING":   true,
		"CONFIRMED": true,
		"CANCELED":  true,
	}
	newStatus = strings.ToUpper(newStatus)
	if !validStatuses[newStatus] {
		return errors.New("invalid status")
	}

	oldStatus := slot.Status
	slot.Status = newStatus

	//  Save slot update
	if err := config.DB.Save(&slot).Error; err != nil {
		return err
	}

	//  Save audit log
	history := models.SlotStatusHistory{
		SlotID:    slot.ID,
		OldStatus: oldStatus,
		NewStatus: newStatus,
		ChangedBy: changedBy,
		ChangedAt: time.Now(),
	}
	return config.DB.Create(&history).Error
}

func GetSlotInfo(slotID int64) (*SlotWithHistoryResponse, error) {
	var slot models.AppointmentSlot
	if err := config.DB.First(&slot, slotID).Error; err != nil {
		return nil, errors.New("slot not found")
	}

	var history []models.SlotStatusHistory
	err := config.DB.
		Where("slot_id = ?", slotID).
		Order("changed_at desc").
		Find(&history).Error
	if err != nil {
		return nil, errors.New("failed to load slot history")
	}

	// ✅ แปลง ChangedAt ให้อยู่ใน timezone Asia/Bangkok
	location, _ := time.LoadLocation("Asia/Bangkok")
	for i := range history {
		history[i].ChangedAt = history[i].ChangedAt.In(location)
	}

	result := &SlotWithHistoryResponse{
		ID:        slot.ID,
		DoctorID:  slot.DoctorID,
		StartTime: slot.StartTime.In(location),
		EndTime:   slot.EndTime.In(location),
		Status:    slot.Status,
		History:   history,
	}
	return result, nil
}