package abstract_factory

type Pricey interface {
	Price() float64
}

type Chair interface {
	Pricey
	IsIoTEnabled() bool
	IsSoft() bool
}

type Dimension struct {
	Length, Width, Weight int
}

type CoffeeTable interface {
	Pricey
	Size() Dimension
	IsFodable() bool
}

type SofaStyle string

const (
	EuropeanStyle SofaStyle = "european"
	AmericanStyle SofaStyle = "american"
)

type Sofa interface {
	Pricey
	Style() SofaStyle
}

type FurnitureFactory interface {
	CreateChair() Chair
	CreateCoffetable() CoffeeTable
	CreateSofa() Sofa
}
