package simpleFactory

type Car interface {
	Name() string
}

type Benz struct {
}

func (b *Benz) Name() string {
	return "benz"
}

type Audi struct {
}

func (a *Audi) Name() string {
	return "audi"
}

type Honda struct {
}

func (h *Honda) Name() string {
	return "honda"
}

func CarFactory(brand string) Car {
	switch brand {
	case "benz":
		return &Benz{}
	case "audi":
		return &Audi{}
	case "honda":
		return &Honda{}
	default:
		return nil
	}
}
