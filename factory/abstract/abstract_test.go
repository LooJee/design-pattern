package abstract

import (
	"testing"
)

func TestCarShop_Brand(t *testing.T) {
	people := NewPeople("小明")

	people.TakeCar(CarShop{}.Brand("benz").ProduceRoadster())
	if people.GetCar() == nil {
		t.Fatalf("benz shouldn't be nil")
	} else if people.GetCar().Name() != "benz_roadster" {
		t.Fatalf("name of benz_roadster is not equal")
	}

	people.TakeCar(CarShop{}.Brand("audi").ProduceRoadster())
	if people.GetCar() == nil {
		t.Fatalf("audi shouldn't be nil")
	} else if people.GetCar().Name() != "audi_roadster" {
		t.Fatalf("name of audi_roadster is not euqal")
	}

	people.TakeCar(CarShop{}.Brand("bmw").ProduceRoadster())
	if people.GetCar() != nil {
		t.Fatalf("bmw should be nil")
	}
}
