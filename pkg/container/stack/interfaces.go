package stack

// IStack Stack interface.
type IStack interface {
	// Push a element to stack top.
	Push(x any)
	// Peek peek stack top element.
	Peek() any
	// Pop a element from stack top.
	Pop() any
}
