package models

import "time"

type Booking struct {
	BookingID    string    `gorm:"primaryKey" json:"booking_id"`
	Restaurant   string    `json:"restaurant_name" binding:"required"`
	CustomerName string    `json:"customer_name" binding:"required"`
	Email        string    `json:"email" binding:"required,email"`
	Phone        string    `json:"phone" binding:"required"`
	TableNumber  int       `json:"table_number" binding:"required"`
	Guests       int       `json:"guests" binding:"required"`
	BookingTime  time.Time `json:"booking_time" binding:"required"`
	Status       string    `json:"status"`
}
