package simple

type ICar interface {
	Name() string
}

type Benz struct {
}

func (b *Benz) Name() string {
	return "benz"
}

func NewBenz() *Benz {
	return &Benz{}
}

type Audi struct {
}

func (a *Audi) Name() string {
	return "audi"
}

func NewAudi() *Audi {
	return &Audi{}
}

/*
simple factory
*/
func CarFactory(brand string) ICar {
	switch brand {
	case "benz":
		return NewBenz()
	case "audi":
		return NewAudi()
	default:
		return nil
	}
}
