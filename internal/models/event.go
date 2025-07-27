package models

import (
	"time"
)

type Event struct {
	ID               string    `json:"id" gorm:"primaryKey;type:varchar(255)"`
	Title            string    `json:"title" gorm:"not null;type:varchar(255)"`
	Description      string    `json:"description" gorm:"type:text"`
	Venue            string    `json:"venue" gorm:"not null;type:varchar(255)"`
	EventDate        time.Time `json:"event_date" gorm:"not null"`
	Price            float64   `json:"price" gorm:"not null;type:decimal(10,2)"`
	TotalTickets     int       `json:"total_tickets" gorm:"not null"`
	AvailableTickets int       `json:"available_tickets" gorm:"not null"`
	ImageUrl         string    `json:"image_url" gorm:"type:varchar(500)"`
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateEventRequest struct {
	Title        string    `json:"title" binding:"required"`
	Description  string    `json:"description"`
	Venue        string    `json:"venue" binding:"required"`
	EventDate    time.Time `json:"event_date" binding:"required"`
	Price        float64   `json:"price" binding:"required,min=0"`
	TotalTickets int       `json:"total_tickets" binding:"required,min=1"`
	ImageUrl     string    `json:"image_url"`
}

type UpdateEventRequest struct {
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Venue            string    `json:"venue"`
	EventDate        time.Time `json:"event_date"`
	Price            float64   `json:"price"`
	TotalTickets     int       `json:"total_tickets"`
	AvailableTickets int       `json:"available_tickets"`
	ImageUrl         string    `json:"image_url"`
}

type UpdateTicketsRequest struct {
	Quantity int `json:"quantity" binding:"required"`
}
