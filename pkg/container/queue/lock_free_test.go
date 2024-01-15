// Package queue
// Create  2023-03-13 19:25:43
package queue

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestQueuePush(t *testing.T) {

	queue := NewLockFreeQueue(4)
	type Obj struct {
		Age int64
	}
	queue.Push(Obj{Age: 999})
	queue.Push(Obj{Age: 999})
	queue.Push(Obj{Age: 999})
	val, err := queue.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Printf("pop result:%+v", val)
}

func TestPointer(t *testing.T) {
	type Obj struct {
		Age int
	}
	var l unsafe.Pointer
	fmt.Println(unsafe.Sizeof(Obj{}))
	ptr := unsafe.Pointer(uintptr(l) + uintptr(unsafe.Sizeof(Obj{})))
	fmt.Printf("ptr:[%p]\n", ptr)
	obj := (*Obj)(ptr)
	obj.Age = 10
	fmt.Println(obj)
}

//func TestQueuePush(t *testing.T) {
//	x := uint64(0x00)
//	atomic.AddUint64(&x, uint64(0x01))
//	fmt.Println(x)
//	atomic.AddUint64(&x, ^uint64(0x00))
//	fmt.Println(x)
//
//	// nil unsafe.Pointer.
//	type Ta structural {
//		ptr unsafe.Pointer
//	}
//	type Obj structural {
//		_ [10]uint64
//	}
//	// 重要.
//	ta := &Ta{}
//	//lp := uintptr(ta.ptr)
//	//fmt.Println(lp)
//	ta.ptr = unsafe.Pointer(uintptr(0) + unsafe.Sizeof(Obj{}))
//	fmt.Printf("%+v", ta)
//}
