package linkedlist

import (
	"github.com/xmopen/golib/pkg/container"
)

// LinkedList go linkedlist.
type LinkedList struct {
	head *node
	tail *node
}

func (l *LinkedList) Push(x any) {
}

func (l *LinkedList) Next() any {
	panic("implement me")
}

func (l *LinkedList) Empty() bool {
	panic("implement me")
}

func (l *LinkedList) Iterators() container.ITerators {
	panic("implement me")
}
