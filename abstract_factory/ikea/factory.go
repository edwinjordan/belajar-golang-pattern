package ikea

//factory
type Ikea struct {
}

func (l *Leifarne) CreateChair() Chair {
	return &Leifarne{}
}

func (l *Leifarne) CreateCoffetable() CoffeTable {
	return &Vittsjo{}
}

func (l *Leifarne) CreateSofa() Sofa {
	return &Hemlingby{}
}
