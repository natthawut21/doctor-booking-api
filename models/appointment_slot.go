package models

import "time"

type AppointmentSlot struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Version   int        `gorm:"column=version" json:"version"`

	DoctorID  uint       `json:"doctor_id"`
	Doctor    Doctor     `gorm:"foreignKey:DoctorID" json:"-"`

	StartTime time.Time  `json:"start_time"`
	EndTime   time.Time  `json:"end_time"`
	Booked    bool       `json:"booked"`
	
}

// กำหนดชื่อ table ให้ตรงกับ Grails
func (AppointmentSlot) TableName() string {
	return "appointment_slot"
}
