package ParkingLotAttendant

import (
	"ParkingLot/ParkingLot"
	"testing"
)

func TestParkingLotCreation(t *testing.T) {
	attendant := NewParkingLotAttendant()
	if attendant == nil {
		t.Errorf("ParkingLotAttendant creation failed")
	}
}

func TestAssignParkingLotByAttendant(t *testing.T) {
	attendant := NewParkingLotAttendant()
	parkingLot, _ := ParkingLot.NewParkingLot(2)
	err := attendant.Assign(parkingLot)
	if err != nil {
		t.Errorf("ParkingLotAttendant Assign failed")
	}
}
