package models

import (
	"time"
)

type Event struct {
	ID               string    `json:"id" db:"id"`
	Title            string    `json:"title" db:"title"`
	Description      string    `json:"description" db:"description"`
	Venue            string    `json:"venue" db:"venue"`
	EventDate        time.Time `json:"event_date" db:"event_date"`
	Price            float64   `json:"price" db:"price"`
	TotalTickets     int       `json:"total_tickets" db:"total_tickets"`
	AvailableTickets int       `json:"available_tickets" db:"available_tickets"`
	ImageUrl         string    `json:"image_url" db:"image_url"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
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
