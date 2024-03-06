package models

import (
	"example.com/rest-api/db"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {

	query := `INSERT INTO USERS(email, password) VALUES(?, ?)`

	preparedQuery, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	result, err := preparedQuery.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

// func GetAllEvents() ([]Event, error) {
// 	query := "SELECT * FROM events;"
// 	rows, err := db.DB.Query(query)

// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()
// 	var events []Event
// 	for rows.Next() {
// 		var event Event
// 		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
// 		fmt.Println(err)
// 		if err != nil {
// 			return nil, err
// 		}

// 		events = append(events, event)
// 	}

// 	return events, nil
// }

// func GetEventById(eventId int64) (*Event, error) {
// 	query := "SELECT * FROM events WHERE id = ?"

// 	row := db.DB.QueryRow(query, eventId)

// 	var event Event
// 	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &event, nil
// }

// func (event Event) Update() error {
// 	query := `UPDATE events
// 	SET name = ?, description = ?, location = ?, dateTime = ?
// 	WHERE id = ?`

// 	preparedQuery, err := db.DB.Prepare(query)
// 	if err != nil {
// 		return err
// 	}
// 	defer preparedQuery.Close()

// 	_, err = preparedQuery.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
// 	return err

// }

// func (event Event) Delete() error {
// 	query := `DELETE from events WHERE id = ?`

// 	preparedQuery, err := db.DB.Prepare(query)
// 	if err != nil {
// 		return err
// 	}
// 	defer preparedQuery.Close()

// 	_, err = preparedQuery.Exec(event.ID)
// 	return err
// }
