package routes

import (
	"bookMyTable/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/bookings")
	{

		api.DELETE("/:id", handlers.DeleteBooking)
	}
}
