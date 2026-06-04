package handlers

import (
	"bookMyTable/config"
	"bookMyTable/models"
	"bookMyTable/utils"

	// "database/sql"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "math/rand"
	// "time"
)

func DeleteBooking(c *gin.Context) {
	id := c.Param("id")

	var booking models.Booking
	if err := config.DB.First(&booking, "booking_id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Booking not found"})
		return
	}

	if err := config.DB.Delete(&booking).Error; err != nil {
		utils.LogError("Delete failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking deleted"})
}
