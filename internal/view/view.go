package view

type View struct {
	Line   int // Start line number
	Column int // Start column
	Width  int
	Height int
}

func (v View) LastColumn() int {
	return v.Column + v.Width - 1
}

func (v View) LastLine() int {
	return v.Line + v.Height - 1
}

func (v View) IsVisible(v2 View) bool {
	if v2.LastColumn() < v.Column ||
		v2.Column > v.LastColumn() ||
		v2.LastLine() < v.Line ||
		v2.Line > v.LastLine() {
		return false
	}
	return true
}

// MinAdjust returns a new View with minimal adjustments so that inner view (iv) is visible.
func (v View) MinAdjust(iv View) View {
	if iv.Width > v.Width || iv.Height > v.Height {
		panic("unimplemented")
	}

	if iv.Column < v.Column {
		v.Column = iv.Column
	} else if iv.LastColumn() > v.LastColumn() {
		v.Column += iv.LastColumn() - v.LastColumn()
	}

	if iv.Line < v.Line {
		v.Line = iv.Line
	} else if iv.LastLine() > v.LastLine() {
		v.Line += iv.LastLine() - v.LastLine()
	}

	return v
}
