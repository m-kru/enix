package mouse

type Event interface {
	isEvent()
	X() int
	Y() int
}

type PrimaryClick struct {
	x int
	y int
}

func (pc PrimaryClick) isEvent() {}
func (pc PrimaryClick) X() int   { return pc.x }
func (pc PrimaryClick) Y() int   { return pc.y }

type DoublePrimaryClick struct {
	x int
	y int
}

func (dpc DoublePrimaryClick) isEvent() {}
func (dpc DoublePrimaryClick) X() int   { return dpc.x }
func (dpc DoublePrimaryClick) Y() int   { return dpc.y }

type TriplePrimaryClick struct {
	x int
	y int
}

func (tpc TriplePrimaryClick) isEvent() {}
func (tpc TriplePrimaryClick) X() int   { return tpc.x }
func (tpc TriplePrimaryClick) Y() int   { return tpc.y }

type PrimaryClickCtrl struct {
	x int
	y int
}

func (pcc PrimaryClickCtrl) isEvent() {}
func (pcc PrimaryClickCtrl) X() int   { return pcc.x }
func (pcc PrimaryClickCtrl) Y() int   { return pcc.y }
