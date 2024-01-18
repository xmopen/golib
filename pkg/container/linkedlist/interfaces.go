package linkedlist

import "github.com/xmopen/golib/pkg/container"

// TODO: ILinkedList 接口设计的不是合理
// 方法名称无法做大见名思意思，双向链表名字挺舒服的.

// ILinkedList  linkedlist interface.
type ILinkedList interface {
	container.IContainer
	Next() any
	// Add 默认添加到队列尾部
	Add(x any)
	// Remove 移除掉队列头部节点
	Remove() any
}

// IDoubleLinkedList interface
type IDoubleLinkedList interface {
	container.IContainer
	PushWithHead(x any)
	PushWithTail(x any)
	PeekFromHead() any
	PeekFromTail() any
	RemoveFromTail() any
	RemoveFromHead() any
}

type node struct {
	item any
	next *node
	prev *node
}
