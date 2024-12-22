package models

import (
	"errors"
	"github.com/npinnaka/goproject/db"
	"log"
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Date        time.Time `json:"date"`
	UserId      int64     `json:"userId"`
}

func (e *Event) Save() error {
	insertStatement := `INSERT INTO events (name, description, location, date, user_id) VALUES (?, ?, ?, ?, ?)`
	result, err := db.DB.Exec(insertStatement, e.Name, e.Description, e.Location, e.Date, e.UserId)
	if err != nil {
		return err
	}
	e.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (e *Event) Update() (*Event, error) {
	updateStatement := `UPDATE events SET name = ?, description = ?, location = ?, date = ?, user_id = ? WHERE id = ?`
	results, err := db.DB.Exec(updateStatement, e.Name, e.Description, e.Location, e.Date, e.UserId, e.ID)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, errors.New("event not found")
	}
	return e, nil
}

func GetAllEvents() ([]Event, error) {
	rows, err := db.DB.Query("SELECT id, name, description, location, date, user_id FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	row := db.DB.QueryRow("SELECT id, name, description, location, date, user_id FROM events where id = ?", id)
	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func DeleteEventById(id int64) (*int64, error) {
	deleteEventQuery := `DELETE FROM events WHERE id = ?`
	result, err := db.DB.Exec(deleteEventQuery, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if affected == 0 {
		return nil, errors.New("Event not found")
	}
	return &affected, nil
}
