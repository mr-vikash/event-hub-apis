package models

import "time"

type Ticket struct {
	ID                 int
	EventId            int
	Name               string
	Description        string
	Price              float64
	Total_Quantity     int
	Remaining_Quantity int
	Sale_Start_Time    time.Time
	Sale_End_Time      time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Status             string
}

func NewTicket(
	EventId int,
	name string,
	description string,
	totalQty int,
	remainingQty int,
	price float64,
	saleStart time.Time,
	saleEnd time.Time,
	CreatedAt time.Time,
	UpdatedAt time.Time,
) Ticket {
	return Ticket{
		EventId:            EventId,
		Name:               name,
		Description:        description,
		Price:              price,
		Total_Quantity:     totalQty,
		Remaining_Quantity: remainingQty, // important logic
		Sale_Start_Time:    saleStart,
		Sale_End_Time:      saleEnd,
		CreatedAt:          CreatedAt,
		UpdatedAt:          UpdatedAt,
		Status:             "active", // default
	}
}
