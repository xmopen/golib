package queue

import (
	"github.com/xmopen/golib/pkg/container"
	"github.com/xmopen/golib/pkg/container/linkedlist"
)

// Queue struct
type Queue struct {
	list linkedlist.IDoubleLinkedList // list 使用双向链表作为队列的容器
}

// New 初始化Queue
func New() IQueue {
	return &Queue{
		list: linkedlist.NewDoubleLinkedList(),
	}
}

// Add 向队头增加元素
// 每次向链表尾部增加数据，每次从头部取出数据用以实现FIFO效果
func (q *Queue) Add(item any) error {
	q.list.PushWithTail(item)
	return nil
}

// Remove 从队列头部移除元素
func (q *Queue) Remove() any {
	return q.list.RemoveFromHead()
}

// Size 返回实际容量大小
func (q *Queue) Size() int {
	return q.list.Size()
}

// Empty 判断Queue是否为空队列
func (q *Queue) Empty() bool {
	return q.list.Empty()
}

// Iterators Queue 迭代器.
// TODO: 先暂时不实现
func (q *Queue) Iterators() container.ITerators {
	panic("implement me")
}

// Poll 移除队列头部元素
func (q *Queue) Poll() any {
	return q.list.RemoveFromHead()
}

// Peek 查看队头元素
func (q *Queue) Peek() any {
	return q.list.PeekFromHead()
}
