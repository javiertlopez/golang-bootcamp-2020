package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/javiertlopez/golang-bootcamp-2020/errorcodes"
	mocks "github.com/javiertlopez/golang-bootcamp-2020/mocks/usecase"
	"github.com/javiertlopez/golang-bootcamp-2020/model"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

const (
	NewStatus      = "NEW"
	OKDescription  = "Wedding"
	BadDescription = "Graduation"
)

func Test_eventController_CreateEvent(t *testing.T) {
	tests := []struct {
		name         string
		expectedCode int
		expectedBody string
		want         model.Event
		errWanted    error
	}{
		{
			"Success",
			201,
			`{"id":"","description":"Wedding","type":"","status":"NEW","created_at":null,"updated_at":null,"event_date":null,"event_location":"","name":"","phone":"","email":""}`,
			model.Event{
				Description: OKDescription,
				Status:      NewStatus,
			},
			nil,
		},
		{
			"Internal error",
			500,
			`{"message":"Internal server error","status":500}`,
			model.Event{},
			errors.New("failed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			events := &mocks.Events{}
			e := &eventController{
				events: events,
			}

			events.On("Create", tt.want).Return(tt.want, tt.errWanted)

			body, err := json.Marshal(tt.want)
			assert.NoError(t, err)

			// Create a request to pass to our handler.
			req, err := http.NewRequest("POST", "/events", bytes.NewBuffer([]byte(body)))
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(e.CreateEvent)

			// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
			// directly and pass in our Request and ResponseRecorder.
			handler.ServeHTTP(rr, req)

			// Check the content type is what we expect.
			expected := "application/json; charset=UTF-8"
			m := rr.Header()
			if contentType := m.Get("Content-Type"); contentType != expected {
				t.Errorf(
					"handler returned wrong content type: got %v want %v",
					contentType,
					expected,
				)
			}

			// Check the status code is what we expect.
			if status := rr.Code; status != tt.expectedCode {
				t.Errorf(
					"handler returned wrong status code: got %v want %v",
					status,
					tt.expectedCode,
				)
			}

			// Check the response body is what we expect.
			if rr.Body.String() != tt.expectedBody {
				t.Errorf(
					"handler returned unexpected body: got %v want %v",
					rr.Body.String(),
					tt.expectedBody,
				)
			}

			events.AssertExpectations(t)
		})
	}

	t.Run("Bad request", func(t *testing.T) {
		e := &eventController{}
		body := `{"description":123,"status":"NEW}`
		expectedCode := 400
		expectedBody := `{"message":"Bad request","status":400}`

		// Create a request to pass to our handler.
		req, err := http.NewRequest("POST", "/events", bytes.NewBuffer([]byte(body)))
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(e.CreateEvent)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the content type is what we expect.
		expected := "application/json; charset=UTF-8"
		m := rr.Header()
		if contentType := m.Get("Content-Type"); contentType != expected {
			t.Errorf(
				"handler returned wrong content type: got %v want %v",
				contentType,
				expected,
			)
		}

		// Check the status code is what we expect.
		if status := rr.Code; status != expectedCode {
			t.Errorf(
				"handler returned wrong status code: got %v want %v",
				status,
				expectedCode,
			)
		}

		// Check the response body is what we expect.
		if rr.Body.String() != expectedBody {
			t.Errorf(
				"handler returned unexpected body: got %v want %v",
				rr.Body.String(),
				expectedBody,
			)
		}
	})
}

func Test_eventController_GetEventByID(t *testing.T) {
	expectedEvent := model.Event{
		ID:          "123",
		Description: OKDescription,
		Status:      NewStatus,
	}

	tests := []struct {
		name         string
		id           string
		expectedCode int
		expectedBody string
		want         model.Event
		errWanted    error
	}{
		{
			"Success",
			"123",
			200,
			`{"id":"123","description":"Wedding","type":"","status":"NEW","created_at":null,"updated_at":null,"event_date":null,"event_location":"","name":"","phone":"","email":""}`,
			expectedEvent,
			nil,
		},
		{
			"Not found",
			"456",
			404,
			`{"message":"Not found","status":404}`,
			model.Event{},
			errorcodes.ErrEventNotFound,
		},
		{
			"Error",
			"789",
			500,
			`{"message":"Internal server error","status":500}`,
			model.Event{},
			errors.New("generic error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			events := &mocks.Events{}
			e := &eventController{
				events: events,
			}

			events.On("GetByID", tt.id).Return(tt.want, tt.errWanted)

			// Create a request to pass to our handler.
			req, err := http.NewRequest("GET", fmt.Sprintf("/events/%s", tt.id), nil)
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/events/{id}", e.GetEventByID)

			// Change to Gorilla Mux router to pass variables
			router.ServeHTTP(rr, req)

			// Check the content type is what we expect.
			expected := "application/json; charset=UTF-8"
			m := rr.Header()
			if contentType := m.Get("Content-Type"); contentType != expected {
				t.Errorf(
					"handler returned wrong content type: got %v want %v",
					contentType,
					expected,
				)
			}

			// Check the status code is what we expect.
			if status := rr.Code; status != tt.expectedCode {
				t.Errorf(
					"handler returned wrong status code: got %v want %v",
					status,
					tt.expectedCode,
				)
			}

			// Check the response body is what we expect.
			if rr.Body.String() != tt.expectedBody {
				t.Errorf(
					"handler returned unexpected body: got %v want %v",
					rr.Body.String(),
					tt.expectedBody,
				)
			}

			events.AssertExpectations(t)
		})
	}
}

func Test_eventController_GetReservations(t *testing.T) {
	arrival := time.Date(2020, 1, 1, 4, 0, 0, 0, time.UTC)
	departure := time.Date(2020, 1, 3, 18, 0, 0, 0, time.UTC)
	reservations := []model.Reservation{
		{
			Adults:    2,
			Minors:    0,
			AdultFee:  7,
			MinorFee:  0,
			Arrival:   &arrival,
			Departure: &departure,
		},
		{
			Adults:    2,
			Minors:    2,
			AdultFee:  7,
			MinorFee:  1,
			Arrival:   &arrival,
			Departure: &departure,
		},
	}

	tests := []struct {
		name         string
		id           string
		expectedCode int
		expectedBody string
		want         []model.Reservation
		errWanted    error
	}{
		{
			"Success",
			"123",
			200,
			`[{"id":"","status":"","plan":"","adults":2,"minors":0,"adult_fee":7,"minor_fee":0,"arrival":"2020-01-01T04:00:00Z","departure":"2020-01-03T18:00:00Z","name":"","phone":"","email":""},{"id":"","status":"","plan":"","adults":2,"minors":2,"adult_fee":7,"minor_fee":1,"arrival":"2020-01-01T04:00:00Z","departure":"2020-01-03T18:00:00Z","name":"","phone":"","email":""}]`,
			reservations,
			nil,
		},
		{
			"Error",
			"123",
			500,
			`{"message":"Internal server error","status":500}`,
			[]model.Reservation{},
			errors.New("failed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			events := &mocks.Events{}
			e := &eventController{
				events: events,
			}

			events.On("GetByID", tt.id).Return(model.Event{}, nil)
			events.On("GetReservations", tt.id).Return(tt.want, tt.errWanted)

			// Create a request to pass to our handler.
			req, err := http.NewRequest("GET", fmt.Sprintf("/events/%s/reservations", tt.id), nil)
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/events/{id}/reservations", e.GetReservations)

			// Change to Gorilla Mux router to pass variables
			router.ServeHTTP(rr, req)

			// Check the content type is what we expect.
			expected := "application/json; charset=UTF-8"
			m := rr.Header()
			if contentType := m.Get("Content-Type"); contentType != expected {
				t.Errorf(
					"handler returned wrong content type: got %v want %v",
					contentType,
					expected,
				)
			}

			// Check the status code is what we expect.
			if status := rr.Code; status != tt.expectedCode {
				t.Errorf(
					"handler returned wrong status code: got %v want %v",
					status,
					tt.expectedCode,
				)
			}

			// Check the response body is what we expect.
			if rr.Body.String() != tt.expectedBody {
				t.Errorf(
					"handler returned unexpected body: got %v want %v",
					rr.Body.String(),
					tt.expectedBody,
				)
			}

			events.AssertExpectations(t)
		})
	}
}
