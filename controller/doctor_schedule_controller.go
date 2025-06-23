package controller

import (
	"doctor-booking-api/models"
	"doctor-booking-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// POST /schedules
func CreateSchedule(c *gin.Context) {
	var schedule models.DoctorSchedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := service.SaveSchedule(schedule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

// PUT /schedules/:id
func UpdateSchedule(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule ID"})
		return
	}

	var schedule models.DoctorSchedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := service.UpdateScheduleByID(uint(id), schedule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

// GET /doctors/:doctorId/schedules
func GetSchedulesByDoctor(c *gin.Context) {
	idParam := c.Param("doctorId")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor ID"})
		return
	}

	schedules, err := service.GetSchedulesByDoctorID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Optional: Sort by DayOfWeek and StartTime if needed here
	c.JSON(http.StatusOK, schedules)
}

// DELETE /schedules?doctorId=1&scheduleId=2
func DeleteSchedule(c *gin.Context) {
	doctorIDStr := c.Query("doctorId")
	scheduleIDStr := c.Query("scheduleId")

	doctorID, err1 := strconv.Atoi(doctorIDStr)
	scheduleID, err2 := strconv.Atoi(scheduleIDStr)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctorId or scheduleId"})
		return
	}

	if err := service.DeleteScheduleByDoctorAndID(uint(doctorID), uint(scheduleID)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule deleted successfully"})
}
