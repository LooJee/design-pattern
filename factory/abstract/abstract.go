package abstract

import "design-pattern/factory/simple"


type ICarFactory interface {
	Produce() simple.ICar
	ProduceRoadster() IRoadster
}

type HondaFactory struct {
}

func (f HondaFactory) Produce() simple.ICar {
	return &simple.Honda{}
}

func (f HondaFactory) ProduceRoadster() IRoadster {
	return &HondaRoadster{}
}

type BenzFactory struct {
}

func (f BenzFactory) Produce() simple.ICar {
	return &simple.Benz{}
}

func (f BenzFactory) ProduceRoadster() IRoadster {
	return &BenzRoadster{}
}

type AudiFactory struct {
}

func (f AudiFactory) Produce() simple.ICar {
	return &simple.Audi{}
}

func (f AudiFactory) ProduceRoadster() IRoadster {
	return &AudiRoadster{}
}

type EmptyFactory struct {

}

func (f EmptyFactory) Produce() simple.ICar {
	return nil
}

func (f EmptyFactory) ProduceRoadster() IRoadster {
	return nil
}

type CarShop struct {

}

func (shop CarShop) Brand(brand string) ICarFactory {
	switch brand {
	case "benz":
		return &BenzFactory{}
	case "audi":
		return &AudiFactory{}
	case "honda":
		return &HondaFactory{}
	default:
		return &EmptyFactory{}
	}
}
