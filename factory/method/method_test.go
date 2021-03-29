package method

import (
	"testing"
)

func TestMethod(t *testing.T) {
	people := NewPeople("小明")

	people.TakeCar(CarShop{}.Brand("benz").Produce())
	if people.GetCar().Name() != "benz" {
		t.Fatalf("name of benz is not equal")
	}

	people.TakeCar(CarShop{}.Brand("audi").Produce())
	if people.GetCar().Name() != "audi" {
		t.Fatalf("name of audi is not equal")
	}
}
