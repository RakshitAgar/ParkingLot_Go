package ParkingLot

import (
	"ParkingLot/Car"
	"ParkingLot/ParkingSlot"
	"ParkingLot/Ticket"
	"errors"
)

type ParkingLot struct {
	IsFull        bool
	LotSize       int
	ParkingSlots  []*ParkingSlot.ParkingSlot
	ParkedSlotMap map[*Ticket.Ticket]*ParkingSlot.ParkingSlot // Using alias 'T'
}

func NewParkingLot(lotSize int) (*ParkingLot, error) {
	if lotSize < 1 {
		return nil, errors.New("lot size must be greater than 0")
	}
	parkingSlots := make([]*ParkingSlot.ParkingSlot, lotSize)
	parkedSlotMap := make(map[*Ticket.Ticket]*ParkingSlot.ParkingSlot)

	for i := 0; i < lotSize; i++ {
		parkingSlots[i] = ParkingSlot.NewParkingSlot(i + 1)
	}

	return &ParkingLot{
		LotSize:       lotSize,
		ParkingSlots:  parkingSlots,
		ParkedSlotMap: parkedSlotMap,
		IsFull:        false,
	}, nil
}

func (lot *ParkingLot) Park(car *Car.Car) (*Ticket.Ticket, error) {

	availableSlot, err := lot.FindNearestAvailableSlot()
	if err != nil {
		return nil, err
	}

	ticket := availableSlot.OccupyWithCar(car)
	lot.ParkedSlotMap[ticket] = availableSlot
	return ticket, nil
}

func (lot *ParkingLot) FindNearestAvailableSlot() (*ParkingSlot.ParkingSlot, error) {
	for _, slot := range lot.ParkingSlots {
		if slot.IsAvailable() {
			return slot, nil
		}
	}
	return nil, errors.New("parking lot is full")
}

func (lot *ParkingLot) CarParkedWithRegistrationNumber(registrationNumber string) (*Car.Car, error) {
	for _, slot := range lot.ParkedSlotMap {
		if slot != nil {
			car, err := slot.GetParkedCar()
			if err == nil && car != nil && car.HasRegistrationNumber(registrationNumber) {
				return car, nil
			}
		}
	}
	return nil, errors.New("car with given registration number not found")
}

func (lot *ParkingLot) UnPark(ticket *Ticket.Ticket) (*Car.Car, error) {
	slot := lot.ParkedSlotMap[ticket]
	if slot == nil {
		return nil, errors.New("invalid Ticket No car found")
	}
	car, _ := slot.GetParkedCar()
	lot.ParkedSlotMap[ticket] = nil
	return car, nil
}
