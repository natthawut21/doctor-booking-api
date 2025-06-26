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
	
	config.DB.AutoMigrate(
		&models.Doctor{},
		&models.User{},
		&models.AppointmentSlot{},
		&models.Appointment{},
		&models.SlotStatusHistory{},
	)
// 	if err := config.DB.AutoMigrate(
// 			&models.Doctor{},
// 	 	&models.User{},
// 	 	&models.AppointmentSlot{},
// 	 	&models.Appointment{},
// 	 	&models.SlotStatusHistory{},); err != nil {
//     log.Fatalf("❌ Failed to migrate: %v", err)
// }

	
	

	r := router.SetupRouter()
	r.Run(":8181")
}
