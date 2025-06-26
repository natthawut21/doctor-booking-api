package models

import "time"

type Doctor struct {
	ID                 int64      `gorm:"primaryKey" json:"id"`
	Version   int    `gorm:"column=version" json:"version"`
	Name               string    `json:"name"`
	Specialty          string    `json:"specialty"`
	SubSpecialty       string    `json:"subSpecialty"`
	Department         string    `json:"department"`
	Phone              string    `json:"phone"`
	Email              string    `json:"email"`
	BankAccountName    string    `json:"bankAccountName"`
	BankAccountNumber  string    `json:"bankAccountNumber"`
	LicenseNumber      string 	 `gorm:"type:varchar(100);unique" json:"licenseNumber"`
	LicenseIssuer      string    `json:"licenseIssuer"`
	LicenseIssuedDate  time.Time `json:"licenseIssuedDate"`
	LicenseExpiryDate  time.Time `json:"licenseExpiryDate"`
}
func (Doctor) TableName() string {
	return "doctor"
}