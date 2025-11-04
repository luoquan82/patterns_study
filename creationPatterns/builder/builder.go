package builder

type House struct {
	Door   string
	Window string
}

type HouseBuilder interface {
	BuildDoor(val string)
	BuildWindow(val string)
	GetHouse() *House
}

type Bao struct {
	house *House
}

func NewBao() *Bao {
	return &Bao{
		house: &House{},
	}
}

func (b *Bao) BuildDoor(val string) {
	b.house.Door = val
}

func (b *Bao) BuildWindow(val string) {
	b.house.Window = val
}

func (b *Bao) GetHouse() *House {
	return b.house
}

type Boss struct {
	builder HouseBuilder
}

func NewBoss(bao *Bao) *Boss {
	return &Boss{
		builder: bao,
	}
}

func (b Boss) GetHouse() *House {
	b.builder.BuildDoor("门")
	b.builder.BuildWindow("窗")
	h := b.builder.GetHouse()
	return h
}
