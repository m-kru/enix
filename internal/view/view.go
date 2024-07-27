package view

type View struct {
	LineNum int // Start line number
	Column  int // Start column
	Width   int
	Height  int
}

func (v View) IsVisible(vis Visible) bool {
	ln := vis.LineNum()
	c := vis.Column()

	if v.LineNum <= ln && ln < v.LineNum+v.Height && v.Column <= c && c < v.Column+v.Width {
		return true
	}

	return false
}

// MinAdjust returns a new View with minimal adjustments so that vis becomes visible.
func (v View) MinAdjust(vis Visible) View {
	c := vis.Column()
	if c < v.Column || v.Column+v.Width <= c {
		v.Column = c
	}

	ln := vis.LineNum()
	if ln < v.LineNum || v.LineNum+v.Height <= ln {
		v.LineNum = ln
	}

	return v
}
