package Ticket

import (
	"github.com/google/uuid"
)

type Ticket struct {
	TicketNumber string
}

func NewTicket() *Ticket {
	return &Ticket{
		TicketNumber: generateTicketNumber(),
	}
}

func generateTicketNumber() string {
	return uuid.New().String()
}
