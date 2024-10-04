package ParkingLot

import (
	ParkingLot2 "ParkingLot/Car"
	"testing"
)

func TestNewParkingLotWithZeroSize(t *testing.T) {
	_, err := NewParkingLot(0)

	if err == nil {
		t.Fatal("Expected an error when creating a ParkingLot with size 0, but got nil")
	}
}

func TestNewParkingLotWithOneSize(t *testing.T) {
	_, err := NewParkingLot(1)
	if err != nil {
		t.Fatal("Not Expected an error when creating a ParkingLot with 1 size")
	}
}

func TestParkingLotParkingOneCar(t *testing.T) {
	firstCar := ParkingLot2.NewCar("red", "AB-12")

	lot, _ := NewParkingLot(5)

	firstTicket, err := lot.Park(firstCar)

	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	if firstTicket == nil {
		t.Fatal("expected ticket, but got nil")
	}

	if lot.ParkedSlotMap[firstTicket] == nil {
		t.Fatalf("expected car to be parked in the lot, but no slot found for ticket")
	}

}

func TestParkingLotParkingTwoCars(t *testing.T) {
	firstCar := ParkingLot2.NewCar("red", "AB-12")
	secondCar := ParkingLot2.NewCar("blue", "AB-12")
	lot, _ := NewParkingLot(5)

	firstTicket, err1 := lot.Park(firstCar)
	secondTicket, err2 := lot.Park(secondCar)
	if err1 != nil {
		t.Fatalf("expected no error, but got %v", err1)
	}

	if firstTicket == nil {
		t.Fatal("expected ticket, but got nil")
	}

	if lot.ParkedSlotMap[firstTicket] == nil {
		t.Fatalf("expected car to be parked in the lot, but no slot found for ticket")
	}

	if err2 != nil {
		t.Fatalf("expected no error, but got %v", err2)
	}

	if secondTicket == nil {
		t.Fatal("expected ticket, but got nil")
	}

	if lot.ParkedSlotMap[secondTicket] == nil {
		t.Fatalf("expected car to be parked in the lot, but no slot found for ticket")
	}

}
