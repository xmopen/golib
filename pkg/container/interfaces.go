package container

// IStack 栈接口.
type IStack interface {
	Push(item any) error
	Pop() (any, error)
	Peek() (any, error)
	Empty() (bool, error)
}

// IContainer 容器基本接口.
type IContainer interface {
	add(item any, index int) error
	get(index int) (any, error)
	del(index int) error
}
