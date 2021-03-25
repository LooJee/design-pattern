package method

import (
	"design-pattern/factory"
	"testing"
)

func TestMethod(t *testing.T) {
	people := factory.NewPeople("小明")

	people.TakeCar(CarShop{}.Brand("honda").Produce())
	if people.GetCar().Name() != "honda" {
		t.Fatalf("name of honda is not equal")
	}

	people.TakeCar(CarShop{}.Brand("benz").Produce())
	if people.GetCar().Name() != "benz" {
		t.Fatalf("name of benz is not equal")
	}

	people.TakeCar(CarShop{}.Brand("audi").Produce())
	if people.GetCar().Name() != "audi" {
		t.Fatalf("name of audi is not equal")
	}
}
