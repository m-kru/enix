package line

func Empty() *Line {
	return &Line{Buf: "", Prev: nil, Next: nil}
}
