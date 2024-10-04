package Car

type Car struct {
	Color              string
	RegistrationNumber string
}

func NewCar(color string, registrationNumber string) *Car {
	return &Car{
		Color:              color,
		RegistrationNumber: registrationNumber,
	}
}

func (car *Car) IsColorValid(color string) bool {
	return car.Color == color
}

func (car *Car) HasRegistrationNumber(registrationNumber string) bool {
	return car.RegistrationNumber == registrationNumber
}
