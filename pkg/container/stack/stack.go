package stack

import (
	"sync"

	"github.com/xmopen/golib/pkg/container"
)

// Stack struct.
type Stack struct {
	container []any // container save element.
	cursor    int   // cursor now.
	lock      *sync.Mutex
}

// New a stack instance.
func New() IStack {
	return &Stack{
		container: make([]any, 0),
		cursor:    -1,
		lock:      &sync.Mutex{},
	}
}

// Push a element to stack top.
func (s *Stack) Push(x any) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.container = append(s.container, x)
	s.cursor++
}

// Pop element from stack top.
// return nil stack is empty.
func (s *Stack) Pop() any {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.Empty() {
		return nil
	}
	topElement := s.container[s.cursor]
	s.container = s.container[0:s.cursor]
	s.cursor--
	return topElement
}

// Peek element from stack top.
func (s *Stack) Peek() any {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.Empty() {
		return nil
	}
	return s.container[s.cursor]
}

// Empty return s is a empty stack.
func (s *Stack) Empty() bool {
	return s.cursor == -1
}

// Iterators return a iterators for stack.
func (s *Stack) Iterators() container.ITerators {
	panic("no implement")
}
