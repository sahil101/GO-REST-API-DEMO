package models

import (
	"fmt"
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserID      int
}

func (e Event) Save() error {

	query := `INSERT INTO EVENTS(name, description, location, dateTime, user_id)
	VALUES(?, ?, ?, ?, ?)`

	preparedQuery, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	result, err := preparedQuery.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events;"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		fmt.Println(err)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(eventId int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"

	row := db.DB.QueryRow(query, eventId)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ? 
	WHERE id = ?`

	preparedQuery, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer preparedQuery.Close()

	_, err = preparedQuery.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err

}

func (event Event) Delete() error {
	query := `DELETE from events WHERE id = ?`

	preparedQuery, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer preparedQuery.Close()

	_, err = preparedQuery.Exec(event.ID)
	return err
}
