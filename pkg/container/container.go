// Package container 容器.
// Create  2023-05-06 21:34:07
package container

// Node 定义结构体中节点.
// 针对树、栈、集合等应该有自己专门TreeNode等,但是应该包含Node数据节点.
type Node struct {
	Val any
}

// IContainer common container interface
// todo Iterators
type IContainer interface {
	Empty() bool
	Iterators() ITerators
}

// ITerators iterators for all container.
type ITerators interface {
}
