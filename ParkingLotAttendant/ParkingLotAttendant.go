package ParkingLotAttendant

import (
	"ParkingLot/Car"
	"ParkingLot/ParkingLot"
	"ParkingLot/Ticket"
	"errors"
)

type ParkingLotAttendant struct {
	AssignedParkingLot []*ParkingLot.ParkingLot
	ParkedCars         []*Car.Car
}

func NewParkingLotAttendant() *ParkingLotAttendant {
	assignedParkingLot := make([]*ParkingLot.ParkingLot, 0)
	parkedCars := make([]*Car.Car, 0)

	return &ParkingLotAttendant{
		AssignedParkingLot: assignedParkingLot,
		ParkedCars:         parkedCars,
	}
}

func (Attendant *ParkingLotAttendant) Assign(parkingLot *ParkingLot.ParkingLot) error {
	for _, lot := range Attendant.AssignedParkingLot {
		if lot == parkingLot {
			return errors.New("parking lot is already assigned")
		}
	}

	Attendant.AssignedParkingLot = append(Attendant.AssignedParkingLot, parkingLot)
	return nil
}

func (Attendant *ParkingLotAttendant) Park(car *Car.Car) (*Ticket.Ticket, error) {
	for _, carParked := range Attendant.ParkedCars {
		if carParked == car {
			return nil, errors.New("car Already Parked")
		}
	}

	for _, lot := range Attendant.AssignedParkingLot {
		if lot == nil {
			return nil, errors.New("no Lot Assigned")
		}
	}

	for _, lot := range Attendant.AssignedParkingLot {
		if !lot.IsFull {

		}
	}
	return nil, nil
}
