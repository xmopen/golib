package linkedlist

import "github.com/xmopen/golib/pkg/container"

// ILinkedList  linkedlist interface.
type ILinkedList interface {
	container.IContainer
	Size() int
	Next() any
	// Add 默认添加到队列尾部
	Add(x any)
	// Remove 移除掉队列头部节点
	Remove()
}

type node struct {
	val  any
	next *node
}
