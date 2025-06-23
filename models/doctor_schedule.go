package models

type DoctorSchedule struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Version   int    `gorm:"column=version" json:"version"`
	DoctorID  uint   `json:"doctor_id"` // FK ไปยัง Doctor
	Doctor    Doctor `gorm:"foreignKey:DoctorID" json:"-"`

	DayOfWeek string `json:"day_of_week"` // "MONDAY", "TUESDAY", ...
	StartTime string `json:"start_time"`  // "09:00"
	EndTime   string `json:"end_time"`    // "17:00"
}
func (DoctorSchedule) TableName() string {
	return "doctor_schedule"
}
