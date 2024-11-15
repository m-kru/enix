package line

func (l *Line) Down() bool {
	if l.Next == nil {
		return false
	}

	nl := l.Next
	nl.Prev = l.Prev
	if l.Prev != nil {
		l.Prev.Next = nl
	}
	if nl.Next != nil {
		nl.Next.Prev = l
	}
	l.Next = nl.Next
	l.Prev = nl
	nl.Next = l

	return true
}

func (l *Line) Up() bool {
	if l.Prev == nil {
		return false
	}

	pl := l.Prev
	pl.Next = l.Next
	if l.Next != nil {
		l.Next.Prev = pl
	}
	if pl.Prev != nil {
		pl.Prev.Next = l
	}
	l.Prev = pl.Prev
	l.Next = pl
	pl.Prev = l

	return true
}
