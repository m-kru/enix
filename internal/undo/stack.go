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
	cap   int
	len   int
	first *Action
	last  *Action
}

func NewStack(cap int) *Stack {
	return &Stack{
		cap:   cap,
		len:   0,
		first: nil,
		last:  nil,
	}
}

func (s *Stack) Clear() {
	s.first = nil
	s.last = nil
	s.len = 0
}

func (s *Stack) Push(act action.Action, curs []*cursor.Cursor, sels []*sel.Selection) {
	if s.cap == 0 {
		return
	}

	if s.len == s.cap {
		first := s.first
		s.first = first.next
		first.next = nil
		if s.first != nil {
			s.first.prev = nil
		} else {
			s.last = nil
		}
		s.len--
	}

	Action := &Action{
		Action:     act,
		Cursors:    curs,
		Selections: sels,
		prev:       nil,
		next:       nil,
	}

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

	act := s.last

	if s.first == s.last {
		s.first = nil
		s.last = nil
	} else {
		s.last = s.last.prev
		s.last.next = nil
	}

	act.prev = nil
	act.next = nil
	s.len--

	return act
}
