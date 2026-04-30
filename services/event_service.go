package services

import (
	"eventhub/models"
	"eventhub/repository"
)

func CreateEvent(name, location string, userID int) error {
	event := models.Event{
		Name:     name,
		Location: location,
		UserID:   userID,
	}

	return repository.CreateEvent(event)
}

func GetEvents() ([]models.Event, error) {
	return repository.GetAllEvents()
}

func JoinEvent(userID, eventID int) error {
	return repository.JoinEvent(userID, eventID)
}
