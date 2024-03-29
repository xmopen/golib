package linkedlist

import (
	"sync"

	"github.com/xmopen/golib/pkg/container"
)

// LinkedList go linkedlist.
type LinkedList struct {
	lock  sync.Locker
	size  int // size LinkedList 持有的元素数量
	first *node
	last  *node
}

// New 初始化LinkedList结构体.
func New() ILinkedList {
	return &LinkedList{
		lock:  &sync.Mutex{},
		size:  0,
		first: nil,
		last:  nil,
	}
}

// Size 链表元素数量
func (l *LinkedList) Size() int {
	return l.size
}

// Add 链表添加节点从尾部
func (l *LinkedList) Add(x any) error {
	l.linkLast(x)
	return nil
}

// linkLast 链表尾部添加数据
func (l *LinkedList) linkLast(x any) {
	element := &node{
		item: x,
		next: nil, // 显示初始化
	}
	l.lock.Lock()
	defer l.lock.Unlock()
	l.size++
	if l.last == nil {
		l.last = element
		l.first = element
		return
	}
	l.last.next = element
	l.last = element
}

// Remove 链表从头部移除节点
func (l *LinkedList) Remove() any {
	removeNode := l.removeFirst()
	if removeNode == nil {
		return nil
	}
	return removeNode.item
}

func (l *LinkedList) removeFirst() *node {
	if l.first == nil {
		return nil
	}
	l.lock.Lock()
	defer l.lock.Unlock()
	temp := l.first
	l.first = temp.next
	l.size--
	return temp
}

// Next 获取链表头部节点，并不会删除
func (l *LinkedList) Next() any {
	firstNode := l.peekFirst()
	if firstNode == nil {
		return nil
	}
	return firstNode.item
}

func (l *LinkedList) peekFirst() *node {
	return l.first
}

func (l *LinkedList) Empty() bool {
	return l.size == 0
}

func (l *LinkedList) Iterators() container.ITerators {
	panic("implement me")
}

// isPositionIndex index在链表内是否有效
// index == l.size 尾部添加
// index < l.size 表示index在size链表内部
func (l *LinkedList) isPositionIndex(index int) bool {
	return index >= 0 || index <= l.size
}
