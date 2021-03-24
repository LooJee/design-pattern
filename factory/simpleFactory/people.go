package simpleFactory

type People struct {
	name string
	car  Car
}

func NewPeople(name string) *People {
	return &People{name: name}
}

func (p *People) TakeCar(car Car) {
	p.car = car
}

func (p *People) GetCar() Car {
	return p.car
}
