package product

import abstract_factory "github.com/imrenagi/design-pattern/abstract_factory"

type VITTSJÖ struct{}

func (v *VITTSJÖ) Price() float64 {
	return 899000
}

func (v *VITTSJÖ) Size() abstract_factory.Dimension {
	return abstract_factory.Dimension{
		Length: 80,
		Width:  70,
		Height: 40,
	}
}

func (v *VITTSJÖ) IsFodable() bool {
	return false
}
