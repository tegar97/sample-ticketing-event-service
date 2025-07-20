package service

import (
	"errors"
	"event-service/internal/models"
	"event-service/internal/repository"
)

type EventService struct {
	eventRepo *repository.EventRepository
}

func NewEventService(eventRepo *repository.EventRepository) *EventService {
	return &EventService{eventRepo: eventRepo}
}

func (s *EventService) GetAllEvents() ([]*models.Event, error) {
	return s.eventRepo.GetAll()
}

func (s *EventService) GetEventByID(id string) (*models.Event, error) {
	if id == "" {
		return nil, errors.New("event ID is required")
	}
	return s.eventRepo.GetByID(id)
}

func (s *EventService) CreateEvent(req *models.CreateEventRequest) (*models.Event, error) {
	event := &models.Event{
		Title:        req.Title,
		Description:  req.Description,
		Venue:        req.Venue,
		EventDate:    req.EventDate,
		Price:        req.Price,
		TotalTickets: req.TotalTickets,
		ImageUrl:     req.ImageUrl,
	}

	err := s.eventRepo.Create(event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (s *EventService) UpdateEvent(id string, req *models.UpdateEventRequest) (*models.Event, error) {
	if id == "" {
		return nil, errors.New("event ID is required")
	}

	existingEvent, err := s.eventRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("event not found")
	}

	if req.Title != "" {
		existingEvent.Title = req.Title
	}
	if req.Description != "" {
		existingEvent.Description = req.Description
	}
	if req.Venue != "" {
		existingEvent.Venue = req.Venue
	}
	if !req.EventDate.IsZero() {
		existingEvent.EventDate = req.EventDate
	}
	if req.Price >= 0 {
		existingEvent.Price = req.Price
	}
	if req.TotalTickets > 0 {
		existingEvent.TotalTickets = req.TotalTickets
	}
	if req.AvailableTickets >= 0 {
		existingEvent.AvailableTickets = req.AvailableTickets
	}
	if req.ImageUrl != "" {
		existingEvent.ImageUrl = req.ImageUrl
	}

	err = s.eventRepo.Update(id, existingEvent)
	if err != nil {
		return nil, err
	}

	return existingEvent, nil
}

func (s *EventService) DeleteEvent(id string) error {
	if id == "" {
		return errors.New("event ID is required")
	}

	_, err := s.eventRepo.GetByID(id)
	if err != nil {
		return errors.New("event not found")
	}

	return s.eventRepo.Delete(id)
}

func (s *EventService) UpdateAvailableTickets(id string, quantity int) error {
	if id == "" {
		return errors.New("event ID is required")
	}

	return s.eventRepo.UpdateAvailableTickets(id, quantity)
}
