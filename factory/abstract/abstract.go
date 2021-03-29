package abstract

import "design-pattern/factory/simple"

type ICarFactory interface {
	Produce() simple.ICar
	ProduceRoadster() IRoadster
}

type BenzFactory struct {
}

func (f BenzFactory) Produce() simple.ICar {
	return simple.NewBenz()
}

func (f BenzFactory) ProduceRoadster() IRoadster {
	return NewBenzRoadster()
}

type AudiFactory struct {
}

func (f AudiFactory) Produce() simple.ICar {
	return simple.NewAudi()
}

func (f AudiFactory) ProduceRoadster() IRoadster {
	return NewAudiRoadster()
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
	default:
		return &EmptyFactory{}
	}
}
