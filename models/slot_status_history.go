package models

import "time"

type SlotStatusHistory struct {
	ID        int64      `gorm:"primaryKey"`
	SlotID    int64      `json:"slot_id"`
	OldStatus string    `json:"old_status"`
	NewStatus string    `json:"new_status"`
	ChangedBy string    `json:"changed_by"` // username หรือ user_id
	ChangedAt time.Time `json:"changed_at"`
	Slot      AppointmentSlot `gorm:"foreignKey:SlotID;references:ID"`
}

func (SlotStatusHistory) TableName() string {
	return "slot_status_history"
}
