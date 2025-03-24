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

type PrimaryClickAlt struct {
	x int
	y int
}

func (pca PrimaryClickAlt) isEvent() {}
func (pca PrimaryClickAlt) X() int   { return pca.x }
func (pca PrimaryClickAlt) Y() int   { return pca.y }

type WheelDown struct {
	x int
	y int
}

func (wd WheelDown) isEvent() {}
func (wd WheelDown) X() int   { return wd.x }
func (wd WheelDown) Y() int   { return wd.y }

type WheelUp struct {
	x int
	y int
}

func (wu WheelUp) isEvent() {}
func (wu WheelUp) X() int   { return wu.x }
func (wu WheelUp) Y() int   { return wu.y }

type WheelLeft struct {
	x int
	y int
}

func (wl WheelLeft) isEvent() {}
func (wl WheelLeft) X() int   { return wl.x }
func (wl WheelLeft) Y() int   { return wl.y }

type WheelRight struct {
	x int
	y int
}

func (wr WheelRight) isEvent() {}
func (wr WheelRight) X() int   { return wr.x }
func (wr WheelRight) Y() int   { return wr.y }
