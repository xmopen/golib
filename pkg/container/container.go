// Package container 容器.
package container

// IContainer common container interface
type IContainer interface {
	Add(item any) error
	Remove() any
	Size() int
	Empty() bool
	Iterators() ITerators
}

// ITerators iterators for all container.
type ITerators interface {
}
