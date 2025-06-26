package router

import (
	"doctor-booking-api/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/doctor", controller.GetDoctors)
	r.POST("/doctor", controller.CreateDoctor)

	// === DoctorSchedule Routes ===
	r.POST("/schedules", controller.CreateSchedule)
	r.PUT("/schedules/:id", controller.UpdateSchedule)
	r.GET("/doctor/:doctorId/schedules", controller.GetSchedulesByDoctor)
	r.DELETE("/schedules", controller.DeleteSchedule)

	r.POST("/slots/generate", controller.GenerateSlots)
	r.GET("/slots/all", controller.ShowAllSlots)
	r.GET("/slots/available", controller.AvailableSlots)
	r.PUT("/slots/:id/status", controller.UpdateSlotStatus)

	r.POST("/appointments/book", controller.BookAppointment)
	r.GET("/appointments", controller.ListAppointments)
	return r
}