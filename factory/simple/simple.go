package simple

type ICar interface {
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

/*
simple factory
*/
func CarFactory(brand string) ICar {
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
