package Ticket

import "testing"

func TestTicketCreation(t *testing.T) {
	ticket := NewTicket()

	if ticket == nil {
		t.Error("Ticket Creation Failed")
	}
}
