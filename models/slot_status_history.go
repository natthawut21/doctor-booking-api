package models

import "time"

type SlotStatusHistory struct {
	ID        uint      `gorm:"primaryKey"`
	SlotID    uint      `json:"slot_id"`
	OldStatus string    `json:"old_status"`
	NewStatus string    `json:"new_status"`
	ChangedBy string    `json:"changed_by"` // username หรือ user_id
	ChangedAt time.Time `json:"changed_at"`
}

func (SlotStatusHistory) TableName() string {
	return "slot_status_history"
}
