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

func (v View) IsVisible(vis Visible) bool {
	ln := vis.LineNum()
	c := vis.Column()

	if v.Line <= ln && ln < v.Line+v.Height && v.Column <= c && c < v.Column+v.Width {
		return true
	}

	return false
}

// MinAdjust returns a new View with minimal adjustments so that vis becomes visible.
func (v View) MinAdjust(vis Visible) View {
	c := vis.Column()
	if c < v.Column {
		v.Column = c
	} else if c > v.LastColumn() {
		v.Column += c - v.LastColumn()
	}

	ln := vis.LineNum()
	if ln < v.Line {
		v.Line = ln
	} else if ln > v.LastLine() {
		v.Line += ln - v.LastLine()
	}

	return v
}
