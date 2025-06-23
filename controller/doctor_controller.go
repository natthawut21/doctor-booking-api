package controller

import (
	"doctor-booking-api/models"
	"doctor-booking-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDoctors(c *gin.Context) {
	doctors, err := service.GetDoctors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, doctors)
}

func CreateDoctor(c *gin.Context) {
	var doctor models.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if doctor.LicenseIssuedDate.IsZero() || doctor.LicenseExpiryDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid license dates"})
		return
	}
	newDoctor, err := service.AddDoctor(doctor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newDoctor)
}
