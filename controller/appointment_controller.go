// controller/appointment_controller.go
package controller

import (
	"doctor-booking-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BookAppointment(c *gin.Context) {
	var req service.BookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	id, err := service.BookAppointment(req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Appointment booked", "id": id})
}

func ListAppointments(c *gin.Context) {
	appointments, err := service.ListAppointments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch appointments"})
		return
	}
	c.JSON(http.StatusOK, appointments)
}
