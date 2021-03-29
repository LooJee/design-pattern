package abstract

import "design-pattern/factory/simple"

type IRoadster interface {
	simple.ICar
	Sport()
}

type BenzRoadster struct {

}

func (r BenzRoadster) Sport() {

}

func (r BenzRoadster) Name() string {
	return "benz_roadster"
}

func NewBenzRoadster() *BenzRoadster {
	return &BenzRoadster{}
}

type AudiRoadster struct {

}

func (r AudiRoadster) Sport() {

}

func (r AudiRoadster) Name() string {
	return "audi_roadster"
}

func NewAudiRoadster() *AudiRoadster {
	return &AudiRoadster{}
}
