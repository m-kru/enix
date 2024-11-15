package undo

import "github.com/m-kru/enix/internal/action"

type action struct {
	action action.Action
	prev   *action
	next   *action
}

type Stack struct {
	len   int
	cap   int
	first *action
	last  *action
}

func NewStack(cap int) *Stack {
	return &Stack{cap: cap}
}

func (s *Stack) Push(act action.Action) {
	a := &action{action: act}

	if s.first == nil {
		s.first = a
		s.last = a
		return
	}

	s.last.next = a
	s.prev = s.last
	s.last = a
}

func (s *Stack) Pop() action.Action {
	a := s.last

	if s.first == s.last {
		s.first = nil
		s.last = nil
	} else {
		s.last = s.last.prev
	}

	return a
}
