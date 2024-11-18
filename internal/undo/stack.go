package undo

import (
	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"
)

type Action struct {
	Action     action.Action
	Cursors    []*cursor.Cursor
	Selections []*sel.Selection
	prev       *Action
	next       *Action
}

type Stack struct {
	len   int
	cap   int
	first *Action
	last  *Action
}

func NewStack(cap int) *Stack {
	return &Stack{cap: cap}
}

func (s *Stack) Push(act action.Action, curs []*cursor.Cursor, sels []*sel.Selection) {
	if s.len == s.cap {
		s.first = s.first.next
		s.len--
	}

	Action := &Action{Action: act, Cursors: curs, Selections: sels}

	if s.first == nil {
		s.first = Action
		s.last = Action
		s.len = 1
		return
	}

	s.last.next = Action
	Action.prev = s.last
	s.last = Action
	s.len++
}

func (s *Stack) Pop() *Action {
	if s.first == nil {
		return nil
	}

	Action := s.last

	if s.first == s.last {
		s.first = nil
		s.last = nil
	} else {
		s.last = s.last.prev
	}

	s.len--

	return Action
}
