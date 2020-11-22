package postgresql

import (
	"fmt"

	"github.com/javiertlopez/golang-bootcamp-2020/model"
	"github.com/javiertlopez/golang-bootcamp-2020/repository"

	"github.com/jmoiron/sqlx"
)

type events struct {
	db *sqlx.DB
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
