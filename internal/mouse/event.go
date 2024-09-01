package mouse

type Event interface {
	isEvent()
}

type PrimaryClick struct {
	X int
	Y int
}

func (pc PrimaryClick) isEvent() {}

type DoublePrimaryClick struct {
	X int
	Y int
}

func (dpc DoublePrimaryClick) isEvent() {}

type TriplePrimaryClick struct {
	X int
	Y int
}

func (tpc TriplePrimaryClick) isEvent() {}
