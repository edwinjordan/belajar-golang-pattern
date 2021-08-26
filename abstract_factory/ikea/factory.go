package ikea

import (
	abstract_factory "github.com/imrenagi/design-pattern/abstract_factory"
	product "github.com/imrenagi/design-pattern/abstract_factory/ikea/product"
)

//factory
type Ikea struct {
}

func (l *Ikea) CreateChair() abstract_factory.Chair {
	return &product.Leifarne{}
}

func (l *Ikea) CreateCoffetable() abstract_factory.CoffeeTable {
	return &product.VITTSJÖ{}
}

func (l *Ikea) CreateSofa() abstract_factory.Sofa {
	return &product.HEMLINGBY{}
}
