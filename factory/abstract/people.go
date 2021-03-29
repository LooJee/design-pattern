package abstract

import "design-pattern/factory/simple"

type People struct {
	name string
	car  simple.ICar
}

func NewPeople(name string) *People {
	return &People{name: name}
}

func (p *People) TakeCar(car simple.ICar) {
	p.car = car
}

func (p *People) GetCar() simple.ICar {
	return p.car
}
