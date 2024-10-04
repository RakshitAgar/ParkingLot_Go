package ParkingSlot

import (
	"ParkingLot/Car"
	"ParkingLot/Ticket"
	"errors"
)

type ParkingSlot struct {
	SlotNumber    int
	ParkedCar     *Car.Car
	slotAvailable bool
}

func NewParkingSlot(slotNumber int) *ParkingSlot {
	return &ParkingSlot{
		SlotNumber:    slotNumber,
		ParkedCar:     nil,
		slotAvailable: true,
	}
}

func (slot *ParkingSlot) IsAvailable() bool {
	return slot.slotAvailable
}

func (slot *ParkingSlot) OccupyWithCar(car *Car.Car) *Ticket.Ticket {
	if slot.IsAvailable() {
		slot.slotAvailable = false
		slot.ParkedCar = car
		return Ticket.NewTicket()
	}
	return nil
}

func (slot *ParkingSlot) GetParkedCar() (*Car.Car, error) {
	if slot.IsAvailable() {
		return nil, errors.New("car Not Found")
	}
	car := slot.ParkedCar
	slot.slotAvailable = true
	slot.ParkedCar = nil
	return car, nil
}

func (slot *ParkingSlot) IsCarOfColor(carColor string) bool {
	return !slot.IsAvailable() && slot.ParkedCar.IsColorValid(carColor)
}

func (slot *ParkingSlot) GetCarByRegistrationNumber(registrationNumber string) (*Car.Car, error) {
	if !slot.IsAvailable() && slot.ParkedCar.HasRegistrationNumber(registrationNumber) {
		return slot.ParkedCar, nil
	}
	return nil, errors.New("car with RegistrationNumber is Not Found")
}
