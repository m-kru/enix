package tab

func (tab *Tab) Change() {
	tab.Delete()
	tab.Insert()
}
