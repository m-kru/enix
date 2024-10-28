package tab

func (tab *Tab) ViewDown() {
	if tab.View.LastLine() >= tab.Lines.Count() {
		return
	}
	tab.View = tab.View.Down()
}

func (tab *Tab) ViewUp() {
	tab.View = tab.View.Up()
}

func (tab *Tab) ViewRight() {
	// - 3 because of:
	// 1. Space between line number and first line character.
	// 2. End of line character,
	// 3. One extra column, it simply looks better.
	lastCol := tab.View.LastColumn() + tab.LineNumWidth() - 3
	if lastCol >= tab.LastColumnIdx() {
		return
	}

	tab.View = tab.View.Right()
}

func (tab *Tab) ViewLeft() {
	tab.View = tab.View.Left()
}
