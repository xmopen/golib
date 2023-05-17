package stack

import (
	"fmt"
	"sync"
	"unsafe"
)

// 先进后出.
// 性能: 应用层零拷贝,直接利用指针.

// 初始化默认值.
const (
	defaultStackSize = 1 << 3

	nodeInvalid = 0
)

// Stack 实现container Stack 接口.
type Stack struct {
	slotSize uintptr
	ptr      unsafe.Pointer

	c     []any
	index int // 当前存储下标.
	cap   int

	lock sync.RWMutex // 尽量避免lock,使用atomic替代,后续和游标一起进行优化.
}

// node 栈中实际存储的结构体.
type node struct {
	d   any
	typ int8 // 0:无效, 1:有效.
}

// New 初始化Stack.
func New(size int) *Stack {
	// 强制size是2的整数倍.
	if size <= 0 {
		//size = defaultStackSize
	}
	stack := &Stack{
		c:        make([]any, 0),
		cap:      size,
		slotSize: unsafe.Sizeof(node{}),
	}
	stack.ptr = unsafe.Pointer(uintptr(unsafe.Pointer(&stack.c)) + uintptr(size)*stack.slotSize)
	// ArbitraryType
	return stack
}

func newNode(item any) *node {
	return &node{
		d: item,
	}
}

// Push 入栈.
func (s *Stack) Push(item any) error {
	if item == nil {
		return nil
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.index == s.cap {
		// debug 模式下可以,运行模式下就可以,很奇怪.
		if err := s.copy(); err != nil {
			return err
		}
	}
	if err := s.add(item, s.index); err != nil {
		return err
	}
	s.index++
	return nil
}

// copy 扩容.
// 每次扩容cap的2倍,当数量大多一定程度时直接按照内存大小来进行扩容.
// TODO: v1先2倍扩容.
func (s *Stack) copy() error {
	size := s.cap * 2
	// 这里扩容,
	if size == 0 {
		size = 1
	}
	s.cap = size
	c := make([]any, 0)
	ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&c)) + uintptr(size)*s.slotSize)
	for i := 0; i < s.index; i++ {
		// 直接将原来指针赋值到现在指针上.
		nod := *(*node)(unsafe.Pointer(uintptr(s.ptr) + uintptr(i)*s.slotSize))
		fmt.Printf("copy idx:[%+v] nod:[%+v]\n", i, nod)
		*(*node)(unsafe.Pointer(uintptr(ptr) + uintptr(i)*s.slotSize)) = nod
	}
	s.ptr = ptr
	return nil
}

// Pop 出栈.
func (s *Stack) Pop(index int) (any, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	//if s.index == 0 {
	//	return nil, nil
	//}
	// get的时候不能以0开始.
	slot, err := s.get(index)
	if err != nil {
		return nil, err
	}
	//s.index--
	return slot.d, nil
}

// Peek 查看栈顶元素.
func (s *Stack) Peek() (any, error) {
	return nil, nil
}

// Empty 查看队列是否为空.
func (s *Stack) Empty() (bool, error) {
	return false, nil
}

func (s *Stack) add(item any, index int) error {
	nod := newNode(item)
	// 把指针nod的值写入到*node中.
	*(*node)(unsafe.Pointer(uintptr(s.ptr) + uintptr(index)*s.slotSize)) = *nod
	fmt.Printf("index:[%+v] any:[%+v]\n", index, item)
	return nil
}

func (s *Stack) get(index int) (*node, error) {
	p := unsafe.Pointer(uintptr(s.ptr) + uintptr(index)*s.slotSize)
	// read: 从指针p开始读取要被转换类型大小个字节.
	val := (*node)(p)
	return val, nil
}

func (s *Stack) del(index int) error {
	return nil
}
