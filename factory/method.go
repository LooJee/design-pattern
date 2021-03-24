package factory

/*
factory method
*/
type ICarFactory interface {
	Produce() ICar
}

type HondaFactory struct {
}

func (f HondaFactory) Produce() ICar {
	return &Honda{}
}

type BenzFactory struct {
}

func (f BenzFactory) Produce() ICar {
	return &Benz{}
}

type AudiFactory struct {
}

func (f AudiFactory) Produce() ICar {
	return &Audi{}
}
