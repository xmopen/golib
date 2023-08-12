// Package queue 无锁队列.
// Create  2023-03-13 19:09:26
package queue

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

// TODO:无锁队列.

// LockFreeQueue 实现无锁队列.
type LockFreeQueue struct {
	queueHeadPtr unsafe.Pointer

	cur  *cursor // 8
	sta  *status
	size uint64
}

// queueItem 队列元素表现形式.
type queueItem struct {
	val *any // 16
}

// cursor 游标,当前队列下个要操作的位置.
type cursor struct {
	// 解决CPU伪共享性能杀手问题: 默认CPU Cache Line 大小为64B, 前后各填充56个字节,用于表示cur独占一个CacheLine.
	_ [6]uint64
	// 高32位用于存储游标,低32位用于存储head.
	cur uint64
	_   [6]uint64
}

const (
// queue =
)

// status 队列访问状态.
type status struct {
	// 解决CPU伪共享性能杀手问题: 默认CPU Cache Line 大小为64B, 前后各填充56个字节,用于表示cur独占一个CacheLine.
	_ [7]uint64
	// 高32位用于存储游标,低32位用于存储当前队列状态.
	sta uint64
	_   [7]uint64
}

// NewLockFreeQueue 构造无锁队列.
func NewLockFreeQueue(size uint64) *LockFreeQueue {
	return &LockFreeQueue{
		size: size,
		cur: &cursor{
			cur: uint64(0x00),
		},
		sta: &status{
			sta: uint64(0x00),
		},
	}
}

// curUnpack 解析游标.
// return cur,status.
func (c *cursor) curUnpack() (uint32, uint32) {
	// c.cur高32位表示当前队列游标.
	return uint32((c.cur >> 32) & (1<<32 - 1)), uint32(c.cur & (1<<32 - 1))
}

// add 游标指针++.
func (c *cursor) add() {
	// 高32位+1.
	// 其实低32位也应该+1.
	atomic.AddUint64(&c.cur, 1<<32)
	// 直接低32位+1.
	atomic.AddUint64(&c.cur, 1)
}

// init 初始化无锁队列内存.
func (l *LockFreeQueue) init() {
	l.queueHeadPtr = unsafe.Pointer(uintptr(0) + uintptr(l.size)*unsafe.Sizeof(queueItem{}))
}

// indexQueue 获取下标cur在数组中的内容.
func (l *LockFreeQueue) indexQueue(cur uint64) *queueItem {
	return (*queueItem)(unsafe.Pointer(uintptr(l.queueHeadPtr) + uintptr(cur)*unsafe.Sizeof(queueItem{})))
}

// tryLockFreeUpdateStatus CAS上锁.
func (l *LockFreeQueue) tryUpdateStatus(oldSta, newSta uint64) {
	// 0x00写可访问.
	// 0x01写不可访问.
	for !atomic.CompareAndSwapUint64(&l.sta.sta, oldSta, newSta) {
		continue
	}
}

// tryReleaseLockFree uint64--.
func (l *LockFreeQueue) tryReleaseLockFree() {
	atomic.AddUint64(&l.cur.cur, ^uint64(0x00))
}

// Push 入队.
// return false 表示入队失败：可能是队列满了.
func (l *LockFreeQueue) Push(val any) bool {
	if val == nil {
		return true
	}
	// 有点类似Java中的Synchronized.
	// TODO: 待优化,应该自选index.
	l.tryUpdateStatus(uint64(0x00), uint64(0x01))
	defer l.tryUpdateStatus(uint64(0x01), uint64(0x00))
	if l.queueHeadPtr == nil {
		l.init()
	}
	cur, head := l.cur.curUnpack()
	// 0111(7) & 0111(7) -> 0111 队满.
	// 环形队列.
	if (head&uint32(l.size-1)) == cur && cur != 0x00 {
		return false
	}

	// 获取对应节点的指针.
	// TODO: 这里貌似有问题.
	qi := l.indexQueue(uint64(cur))
	//if qi.val != nil {
	//	panic(fmt.Sprintf("queue index conflict:[%+v]", cur))
	//}

	*(qi.val) = &val
	// 更新游标和head头.
	l.cur.add()
	return true
}

var ErrorQueueEmpty = fmt.Errorf("empty queue")

// Pop 出队.
func (l *LockFreeQueue) Pop() (any, error) {
	if l.queueHeadPtr == nil {
		panic(fmt.Sprintf("pop a invalid queue"))
	}
	l.tryUpdateStatus(uint64(0x00), uint64(0x01))
	defer l.tryUpdateStatus(uint64(0x01), uint64(0x00))
	_, head := l.cur.curUnpack()
	qi := l.indexQueue(uint64(head))
	if qi.val == nil {
		// TODO: 后续通过信号量机制进行Wait.
		return nil, ErrorQueueEmpty
	}
	atomic.AddUint64(&l.cur.cur, uint64(0x01))
	return *qi.val, nil
}

// Len 返回队列当前元素长度.
func (l *LockFreeQueue) Len() int64 {
	return 0
}

// String Stringer.
func (l *LockFreeQueue) String() string {
	return ""
}
