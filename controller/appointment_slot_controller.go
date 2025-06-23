package controller

import (
	"doctor-booking-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// POST /appointments/generate?doctorId=1&date=2025-06-20
func GenerateSlots(c *gin.Context) {
	doctorIDStr := c.Query("doctorId")
	date := c.Query("date")

	doctorID, err := strconv.Atoi(doctorIDStr)
	if err != nil || date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid doctorId or date"})
		return
	}

	slots, err := service.GenerateSlots(uint(doctorID), date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, slots)
}
func ShowAllSlots(c *gin.Context) {
	doctorIDStr := c.Query("doctorId")
	dateStr := c.Query("date")

	doctorID, err := strconv.ParseUint(doctorIDStr, 10, 64)
	if err != nil || dateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid doctorId or date"})
		return
	}

	slots, err := service.ShowAllSlots(uint(doctorID), dateStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, slots)
}

func AvailableSlots(c *gin.Context) {
	doctorIDStr := c.Query("doctorId")
	dateStr := c.Query("date")

	doctorID, err := strconv.ParseUint(doctorIDStr, 10, 64)
	if err != nil || dateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid doctorId or date"})
		return
	}

	slots, err := service.AvailableSlots(uint(doctorID), dateStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, slots)
}

