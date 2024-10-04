package Car

import "testing"

func TestCarCreation(t *testing.T) {
	car := NewCar("Red", "Up81")
	if car == nil {
		t.Fatal("Car creation failed")
	}
}

func TestCarWithSameColor(t *testing.T) {
	car := NewCar("Red", "Up81")

	isColorValid := car.IsColorValid("Red")

	if isColorValid != true {
		t.Fatal("Error occur while validating the Color")
	}
}

func TestCarWithDifferentColor(t *testing.T) {
	car := NewCar("Red", "Up81")

	isColorValid := car.IsColorValid("Blue")

	if isColorValid != false {
		t.Fatal("Error occur while validating the Color")
	}
}

func TestCarWithSameRegistrationNumber(t *testing.T) {
	car := NewCar("Red", "Up81")

	isColorValid := car.HasRegistrationNumber("Up81")

	if isColorValid != true {
		t.Fatal("Error occur while validating the Color")
	}
}

func TestCarWithDifferentRegistrationNumber(t *testing.T) {
	car := NewCar("Red", "Up81")

	isColorValid := car.HasRegistrationNumber("Up82")

	if isColorValid != false {
		t.Fatal("Error occur while validating the Color")
	}
}
