package linkedlist

// ILinkedList  linkedlist interface.
type ILinkedList interface {
	Push(x any)
	Next() any
}

type node struct {
	val any
}
