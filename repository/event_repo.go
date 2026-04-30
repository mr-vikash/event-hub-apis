package repository

import (
	"eventhub/config"
	"eventhub/models"
)

func CreateEvent(event models.Event) error {
	query := "INSERT INTO events(name, location, user_id) VALUES(?, ?, ?)"
	_, err := config.DB.Exec(query, event.Name, event.Location, event.UserID)
	return err
}

func GetAllEvents() ([]models.Event, error) {
	rows, err := config.DB.Query("SELECT id, name, location, user_id FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event

	for rows.Next() {
		var e models.Event
		rows.Scan(&e.ID, &e.Name, &e.Location, &e.UserID)
		events = append(events, e)
	}

	return events, nil
}

func JoinEvent(userID, eventID int) error {
	query := "INSERT INTO event_participants(user_id, event_id) VALUES(?, ?)"
	_, err := config.DB.Exec(query, userID, eventID)
	return err
}
