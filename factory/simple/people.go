package simple

type People struct {
	name string
	car  ICar
}

func NewPeople(name string) *People {
	return &People{name: name}
}

func (p *People) TakeCar(car ICar) {
	p.car = car
}

func (p *People) GetCar() ICar {
	return p.car
}
