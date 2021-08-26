package product

type Beanbag struct {
	price         float64
	Softnesslevel int
}

func (b *Beanbag) Price() float64 {
	return 1190000
}

func (b *Beanbag) IsIoTEnabled() bool {
	return false
}

func (b *Beanbag) IsSoft() bool {
	return b.Softnesslevel > 10
}
