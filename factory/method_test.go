package factory

import "testing"

func TestMethod(t *testing.T) {
	people := NewPeople("小明")

	people.TakeCarFromFactory(HondaFactory{})
	if people.GetCar().Name() != "honda" {
		t.Fatalf("name of honda is not equal")
	}

	people.TakeCarFromFactory(BenzFactory{})
	if people.GetCar().Name() != "benz" {
		t.Fatalf("name of benz is not equal")
	}

	people.TakeCarFromFactory(AudiFactory{})
	if people.GetCar().Name() != "audi" {
		t.Fatalf("name of audi is not equal")
	}
}
