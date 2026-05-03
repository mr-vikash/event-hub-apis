package services

import (
	"eventhub/models"
	"eventhub/repository"
	"time"
)

func CreateTicket(event_id int, name, description string, total_quantity, remaining_quantity int, price float64, sale_start_time, sale_end_time time.Time) error {
	ticket := models.NewTicket(
		event_id,
		name,
		description,
		total_quantity,
		remaining_quantity,
		price,
		sale_start_time,
		sale_end_time,
		time.Now(),
		time.Now(),
	)

	err := repository.CreateTicket(ticket)

	if err != nil {
		return err
	}

	return nil
}
