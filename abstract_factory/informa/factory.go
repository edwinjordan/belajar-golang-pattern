package informa

import (
	abstract_factory "github.com/imrenagi/design-pattern/abstract_factory"
	product "github.com/imrenagi/design-pattern/abstract_factory/informa/product"
)

//factory
type Informa struct {
}

func (l *Informa) CreateChair() abstract_factory.Chair {
	//return &product.BeanBag{}

	return &product.BeanBag{
		// price : product.price,
		// SoftnessLevel: product.SoftnessLevel,
	}
}

func (l *Informa) CreateCoffetable() abstract_factory.CoffeeTable {
	return nil
}

func (l *Informa) CreateSofa() abstract_factory.Sofa {
	return nil
}
