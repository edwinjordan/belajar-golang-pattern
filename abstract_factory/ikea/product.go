package ikea

type Leifarne struct{}

func (l *Leifarne) Price() float64 {
	return 1095000
}

func (l *Leifarne) IsIoTEnabled() bool {
	return false
}

func (l *Leifarne) IsSoft() bool {
	return false
}

type Vittsjo struct{}

func (v *Vittsjo) Price() float64 {
	return 899000
}

func (v *Vittsjo) Size() D {
	return Dimension{
		Length: 80,
		Width:  70,
		Weight: 40,
	}
}

func (v *Vittsjo) IsFodable() bool {
	return false
}

type Hemlingby struct{}

func (h *Hemlingby) Style() SofaStyle {
	return EuropeanStyle
}

func (h *Hemlingby) Price() float64 {
	return 1795000
}
