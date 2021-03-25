package method

import "design-pattern/factory/simple"

/*
factory method
*/
type ICarFactory interface {
	Produce() simple.ICar
}

type HondaFactory struct {
}

func (f HondaFactory) Produce() simple.ICar {
	return simple.NewHonda()
}

type BenzFactory struct {
}

func (f BenzFactory) Produce() simple.ICar {
	return simple.NewBenz()
}

type AudiFactory struct {
}

func (f AudiFactory) Produce() simple.ICar {
	return simple.NewAudi()
}

type EmptyFactory struct {

}

func (f EmptyFactory) Produce() simple.ICar {
	return nil
}

type CarShop struct {

}

func (shop CarShop) Brand(brand string) ICarFactory {
	switch brand {
	case "honda":
		return HondaFactory{}
	case "benz":
		return BenzFactory{}
	case "audi":
		return AudiFactory{}
	default:
		return EmptyFactory{}
	}
}
