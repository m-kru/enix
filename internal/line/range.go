package line

type Range struct {
	Lower int
	Upper int
}

func (r Range) Width() int {
	return r.Upper - r.Lower + 1
}

func (r Range) Within(idx int) bool {
	if r.Lower <= idx && idx <= r.Upper {
		return true
	}
	return false
}
