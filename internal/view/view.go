package view

type View struct {
	Line   int // Start line number
	Column int // Start column
	Height int
	Width  int
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

// Intersection returns intersection of two views.
// It is users responsibility to first check if two views
// intersect using the IsVisible function.
func (v View) Intersection(v2 View) View {
	iv := View{}

	if v.Line <= v2.Line {
		iv.Line = v2.Line
	} else {
		iv.Line = v.Line
	}
	lastLine := v.LastLine()
	if v.LastLine() > v2.LastLine() {
		lastLine = v2.LastLine()
	}
	iv.Height = lastLine - iv.Line + 1

	if v.Column <= v2.Column {
		iv.Column = v2.Column
	} else {
		iv.Column = v.Column
	}
	lastCol := v.LastColumn()
	if v.LastColumn() > v2.LastColumn() {
		lastCol = v2.LastColumn()
	}
	iv.Width = lastCol - iv.Column + 1

	return iv
}

// MinAdjust returns a new View with minimal adjustments so that inner view (iv) is visible.
func (v View) MinAdjust(iv View) View {
	// Under what circumstances inner view can span greater area than view?
	// It is not yet clear.
	if iv.Width > v.Width || iv.Height > v.Height {
		return iv.MinAdjust(v)
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

func (v View) Down() View {
	v.Line += 1
	return v
}

func (v View) Up() View {
	v.Line -= 1
	if v.Line < 1 {
		v.Line = 1
	}
	return v
}

func (v View) Right() View {
	v.Column += 1
	return v
}

func (v View) Left() View {
	v.Column -= 1
	if v.Column < 1 {
		v.Column = 1
	}
	return v
}
