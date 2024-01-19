package linkedlist

import (
	"sync"

	"github.com/xmopen/golib/pkg/container"
)

// DoubleLinkedList 双向链表
type DoubleLinkedList struct {
	lock sync.Locker
	size int
	head *node
	tail *node
}

// NewDoubleLinkedList 初始化双向链表
func NewDoubleLinkedList() IDoubleLinkedList {
	return &DoubleLinkedList{
		lock: &sync.Mutex{},
		size: 0,
		head: nil,
		tail: nil,
	}
}

// Add DoubleLinkedList从尾部插入节点
func (d *DoubleLinkedList) Add(item any) error {
	d.PushWithTail(item)
	return nil
}

// PushWithHead Push node to head
func (d *DoubleLinkedList) PushWithHead(x any) {
	nowNode := &node{
		item: x,
		next: d.head,
		prev: nil,
	}
	d.lock.Lock()
	defer d.lock.Unlock()
	d.size++
	d.head = nowNode
	if d.tail == nil {
		d.tail = nowNode
		return
	}
	// 更新旧的头节点的前驱节点
	d.head.next.prev = d.head
}

// PushWithTail 向链表尾节点插入数据
func (d *DoubleLinkedList) PushWithTail(x any) {
	nowNode := &node{
		item: x,
		next: nil,
		prev: d.tail,
	}
	d.lock.Lock()
	d.lock.Unlock()
	d.size++
	d.tail = nowNode
	if d.head == nil {
		d.head = nowNode
		return
	}
	// 更新旧的尾部节点的后驱节点
	d.tail.prev.next = d.tail
}

// Remove DoubleLinkedList 从尾部移除掉某个节点
func (d *DoubleLinkedList) Remove() any {
	return d.RemoveFromTail()
}

// RemoveFromHead 双向链表移除头部节点
func (d *DoubleLinkedList) RemoveFromHead() any {
	if d.head == nil {
		return nil
	}
	d.lock.Lock()
	d.lock.Unlock()
	d.size--
	tempHead := d.head
	d.head = d.head.next
	// 此时DoubleLinkedList只有一个节点
	if d.head == nil {
		return tempHead.item
	}
	d.head.prev = nil
	return tempHead.item
}

// RemoveFromTail 双向链表移除尾部节点
func (d *DoubleLinkedList) RemoveFromTail() any {
	if d.tail == nil {
		return nil
	}
	d.lock.Lock()
	defer d.lock.Unlock()
	d.size--
	tempTail := d.tail
	d.tail = d.tail.prev
	if d.tail == nil {
		return tempTail.item
	}
	d.tail.next = nil
	return tempTail.item
}

// PeekFromHead Peek Head
func (d *DoubleLinkedList) PeekFromHead() any {
	if d.Empty() {
		return nil
	}
	return d.head.item
}

// PeekFromTail Peek Tail
func (d *DoubleLinkedList) PeekFromTail() any {
	if d.Empty() {
		return nil
	}
	return d.tail.item
}

// Empty DoubleLinkedList is empty list
func (d *DoubleLinkedList) Empty() bool {
	return d.size == 0
}

// Size DoubleLinkedList size
func (d *DoubleLinkedList) Size() int {
	return d.size
}

// Iterators 双向链表迭代器，TODO：暂未实现
func (d *DoubleLinkedList) Iterators() container.ITerators {
	panic("implement me")
}
