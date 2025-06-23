// models/appointment.go
package models

import "time"

type Appointment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Version   int            `gorm:"default:0" json:"version"`
	DoctorID  uint           `json:"doctor_id"`
	PatientID uint           `json:"patient_id"`
	SlotID    uint           `gorm:"unique" json:"slot_id"` // ห้ามซ้ำ
	CreatedAt time.Time      `json:"created_at"`

	Doctor    Doctor         `gorm:"foreignKey:DoctorID"`
	Patient   User           `gorm:"foreignKey:PatientID"`
	Slot      AppointmentSlot `gorm:"foreignKey:SlotID"`
}
func (Appointment) TableName() string {
	return "appointment"
}