package simple

import (
	"testing"
)

func TestCarFactory(t *testing.T) {
	people := NewPeople("小明")

	people.TakeCar(CarFactory("bmw"))
	if people.GetCar() != nil {
		t.Fatalf("bmw is not a valid car")
	}

	people.TakeCar(CarFactory("audi"))
	if people.GetCar() == nil {
		t.Fatalf("audi should be a valid car")
	} else if people.GetCar().Name() != "audi" {
		t.Fatalf("name of audi is not equal")
	}

	people.TakeCar(CarFactory("benz"))
	if people.GetCar() == nil {
		t.Fatalf("benz should be a valid car")
	} else if people.GetCar().Name() != "benz" {
		t.Fatalf("name of benz is not equal")
	}
}
