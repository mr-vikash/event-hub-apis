package repository

import (
	"eventhub/config"
	"eventhub/models"
)

func CreateTicket(ticket models.Ticket) error {
	query := `
INSERT INTO tickets (
    event_id,
    name,
    description,
    price,
    total_quantity,
    remaining_quantity,
    sale_start_time,
    sale_end_time,
    status,
    created_at,
    updated_at
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

	_, err := config.DB.Exec(
		query,
		ticket.EventId,
		ticket.Name,
		ticket.Description,
		ticket.Price,
		ticket.Total_Quantity,
		ticket.Remaining_Quantity,
		ticket.Sale_Start_Time,
		ticket.Sale_End_Time,
		ticket.Status,
		ticket.CreatedAt,
		ticket.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
