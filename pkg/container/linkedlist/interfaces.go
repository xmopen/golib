package linkedlist

import "github.com/xmopen/golib/pkg/container"

// TODO: ILinkedList 接口设计的不是合理
// 方法名称无法做大见名思意思，双向链表名字挺舒服的.

// TODO: 支持异步 lock.
// 低效：每一个操作都进行Lock,性能低下
// 中等：Write时加Lock,
// 高效：无锁CAS

// ILinkedList  linkedlist interface.
type ILinkedList interface {
	container.IContainer
	Next() any
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
	RemoveWithValue(item any) any
}

type node struct {
	item any
	next *node
	prev *node
}
