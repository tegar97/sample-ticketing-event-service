package repository

import (
	"errors"
	"event-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) GetAll() ([]*models.Event, error) {
	var events []*models.Event
	err := r.db.Order("event_date ASC").Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) GetByID(id string) (*models.Event, error) {
	event := &models.Event{}
	err := r.db.Where("id = ?", id).First(event).Error
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (r *EventRepository) Create(event *models.Event) error {
	event.ID = uuid.New().String()
	event.AvailableTickets = event.TotalTickets

	return r.db.Create(event).Error
}

func (r *EventRepository) Update(id string, event *models.Event) error {
	return r.db.Where("id = ?", id).Updates(event).Error
}

func (r *EventRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&models.Event{}).Error
}

func (r *EventRepository) UpdateAvailableTickets(id string, quantity int) error {
	result := r.db.Model(&models.Event{}).
		Where("id = ? AND available_tickets + ? >= 0 AND available_tickets + ? <= total_tickets", id, quantity, quantity).
		Update("available_tickets", gorm.Expr("available_tickets + ?", quantity))

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no rows affected: invalid ticket quantity or event not found")
	}

	return nil
}
