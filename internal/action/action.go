package action

type Action interface {
	isAction()
	Reverse() Action
}

type Actions []Action

func (as Actions) isAction() {}

func (as Actions) Reverse() Action {
	ras := make(Actions, 0, len(as))

	for i := len(as) - 1; i >= 0; i-- {
		a := as[i]
		ras = append(ras, a.Reverse())
	}

	return ras
}
