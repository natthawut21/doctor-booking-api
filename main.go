package main

import (
	"doctor-booking-api/config"
	"doctor-booking-api/models"
	"doctor-booking-api/router"
	"time"
)

func main() {

	loc, _ := time.LoadLocation("Asia/Bangkok")
    time.Local = loc


	config.ConnectDatabase()
	// ✅ AutoMigrate ตารางทั้งหมดที่ใช้ในระบบ
	config.DB.AutoMigrate(
		&models.Doctor{},
		&models.User{},
		&models.AppointmentSlot{},
		&models.Appointment{},
		&models.SlotStatusHistory{},
	)

	r := router.SetupRouter()
	r.Run(":8181")
}
