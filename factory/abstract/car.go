package abstract

import "design-pattern/factory/simple"

type IRoadster interface {
	simple.ICar
	Sport()
}

type HondaRoadster struct {
}

func (r HondaRoadster) Sport() {
}

func (r HondaRoadster) Name() string {
	return "honda_roadster"
}

type BenzRoadster struct {

}

func (r BenzRoadster) Sport() {

}

func (r BenzRoadster) Name() string {
	return "benz_roadster"
}

type AudiRoadster struct {

}

func (r AudiRoadster) Sport() {

}

func (r AudiRoadster) Name() string {
	return "audi_roadster"
}