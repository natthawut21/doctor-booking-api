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

	slots, err := service.GenerateSlots(int64(doctorID), date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, slots)
}
func ShowAllSlots(c *gin.Context) {
	doctorIDStr := c.Query("doctorId")
	dateStr := c.Query("date")

	doctorID, err := strconv.ParseInt(doctorIDStr, 10, 64)
	if err != nil || dateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid doctorId or date"})
		return
	}

	slots, err := service.ShowAllSlots(int64(doctorID), dateStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, slots)
}

func AvailableSlots(c *gin.Context) {
	doctorIDStr := c.Query("doctorId")
	dateStr := c.Query("date")

	doctorID, err := strconv.ParseInt(doctorIDStr, 10, 64)
	if err != nil || dateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid doctorId or date"})
		return
	}

	slots, err := service.AvailableSlots(int64(doctorID), dateStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, slots)
}
func BookedSlots(c *gin.Context) {
	doctorIDStr := c.Query("doctorId")
	dateStr := c.Query("date")

	doctorID, err := strconv.ParseInt(doctorIDStr, 10, 64)
	if err != nil || dateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid doctorId or date"})
		return
	}

	slots, err := service.BookedSlots(int64(doctorID), dateStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, slots)
}


type UpdateStatusRequest struct {
	Status string `json:"status"`
	ChangedBy string `json:"changed_by"` // เช่น username
}


func UpdateSlotStatus(c *gin.Context) {
	idParam := c.Param("id")
	slotID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid slot ID"})
		return
	}

	var req UpdateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = service.UpdateSlotStatus(int64(slotID), req.Status, req.ChangedBy)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Slot status updated"})
}

func GetSlotInfo(c *gin.Context) {
	idStr := c.Param("id")
	slotID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid slot ID"})
		return
	}

	result, err := service.GetSlotInfo(slotID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
