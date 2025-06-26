// models/user.go
package models

type User struct {
	ID       int64   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Name     string `json:"name"`
}
func (User) TableName() string {
	return "user"
}