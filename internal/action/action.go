package action

type Action interface {
	isAction()
	Reverse() Action
}
