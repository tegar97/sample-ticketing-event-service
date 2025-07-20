package repository

import (
	"database/sql"
	"event-service/internal/models"

	"github.com/google/uuid"
)

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) GetAll() ([]*models.Event, error) {
	query := `
        SELECT id, title, description, venue, event_date, price, total_tickets, available_tickets, image_url, created_at, updated_at
        FROM events
        ORDER BY event_date ASC
    `

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*models.Event
	for rows.Next() {
		event := &models.Event{}
		err := rows.Scan(
			&event.ID, &event.Title, &event.Description, &event.Venue,
			&event.EventDate, &event.Price, &event.TotalTickets, &event.AvailableTickets,
			&event.ImageUrl, &event.CreatedAt, &event.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (r *EventRepository) GetByID(id string) (*models.Event, error) {
	event := &models.Event{}

	query := `
        SELECT id, title, description, venue, event_date, price, total_tickets, available_tickets, image_url, created_at, updated_at
        FROM events
        WHERE id = $1
    `

	err := r.db.QueryRow(query, id).Scan(
		&event.ID, &event.Title, &event.Description, &event.Venue,
		&event.EventDate, &event.Price, &event.TotalTickets, &event.AvailableTickets,
		&event.ImageUrl, &event.CreatedAt, &event.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (r *EventRepository) Create(event *models.Event) error {
	event.ID = uuid.New().String()
	event.AvailableTickets = event.TotalTickets

	query := `
        INSERT INTO events (id, title, description, venue, event_date, price, total_tickets, available_tickets, image_url)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING created_at, updated_at
    `

	err := r.db.QueryRow(query, event.ID, event.Title, event.Description, event.Venue,
		event.EventDate, event.Price, event.TotalTickets, event.AvailableTickets, event.ImageUrl).
		Scan(&event.CreatedAt, &event.UpdatedAt)

	return err
}

func (r *EventRepository) Update(id string, event *models.Event) error {
	query := `
        UPDATE events 
        SET title = $1, description = $2, venue = $3, event_date = $4, price = $5, 
            total_tickets = $6, available_tickets = $7, image_url = $8, updated_at = CURRENT_TIMESTAMP
        WHERE id = $9
        RETURNING updated_at
    `

	err := r.db.QueryRow(query, event.Title, event.Description, event.Venue,
		event.EventDate, event.Price, event.TotalTickets, event.AvailableTickets, event.ImageUrl, id).
		Scan(&event.UpdatedAt)

	return err
}

func (r *EventRepository) Delete(id string) error {
	query := `DELETE FROM events WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *EventRepository) UpdateAvailableTickets(id string, quantity int) error {
	query := `
        UPDATE events 
        SET available_tickets = available_tickets + $1, updated_at = CURRENT_TIMESTAMP
        WHERE id = $2 AND available_tickets + $1 >= 0 AND available_tickets + $1 <= total_tickets
    `

	result, err := r.db.Exec(query, quantity, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
