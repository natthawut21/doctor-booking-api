package models

import "time"

type AppointmentSlot struct {
	ID        int64       `gorm:"primaryKey" json:"id"`
	Version   int        `gorm:"column=version" json:"version"`

	DoctorID  int64       `json:"doctor_id"`
	Doctor    Doctor     `gorm:"foreignKey:DoctorID" json:"-"`

	StartTime time.Time  `json:"start_time"`
	EndTime   time.Time  `json:"end_time"`
	
	// จาก boolean → เป็น enum status แทน
	//Status string `gorm:"type:enum('AVAILABLE','PENDING','CONFIRMED','CANCELED'); default:'AVAILABLE'" json:"status"`
	Status string `gorm:"default:'AVAILABLE'" json:"status"`
}

// กำหนดชื่อ table ให้ตรงกับ Grails
func (AppointmentSlot) TableName() string {
	return "appointment_slot"
}
