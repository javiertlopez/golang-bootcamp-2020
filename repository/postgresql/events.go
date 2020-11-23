package postgresql

import (
	"fmt"
	"time"

	"github.com/javiertlopez/golang-bootcamp-2020/model"
	"github.com/javiertlopez/golang-bootcamp-2020/repository"

	"github.com/jmoiron/sqlx"
)

type events struct {
	db *sqlx.DB
}

type event struct {
	ID            string     `db:"id"`
	Description   string     `db:"description"`
	Type          string     `db:"type"` // how can I predefine values here?
	Status        string     `db:"status"`
	CreatedAt     *time.Time `db:"created_at"`
	UpdatedAt     *time.Time `db:"updated_at"`
	EventDate     *time.Time `db:"event_date"`     // should I drop 'event_'?
	EventLocation string     `db:"event_location"` // should I drop 'event_'?

	// Customer information
	Name  string `db:"name"`
	Phone string `db:"phone"`
	Email string `db:"email"`
}

// NewEventsRepo returns the EventRepository implementation
func NewEventsRepo(db *sqlx.DB) repository.EventRepository {
	return &events{
		db,
	}
}

func (e *events) Create(event model.Event) (model.Event, error) {
	return event, fmt.Errorf("not implemented")
}

func (e *events) GetByID(id string) (model.Event, error) {
	return model.Event{}, fmt.Errorf("not implemented")

}
func (e *events) GetAll() ([]model.Event, error) {
	return nil, fmt.Errorf("not implemented")

}
func (e *events) Update(event model.Event) (model.Event, error) {
	return event, fmt.Errorf("not implemented")

}
func (e *events) Delete(id string) error {
	return fmt.Errorf("not implemented")
}
