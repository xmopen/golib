package main

import (
	"context"
	"fmt"
	"sync"
	"time"
	"unsafe"

	"golang.org/x/time/rate"
)

type Object struct {
	//Name string
	A any // 占用16个字节.
}

func testSyncPool() {
	poll := &sync.Pool{
		New: func() interface{} {
			return &Object{}
		},
	}
	for i := 0; i < 38; i++ {
		go func() {
			poll.Put(&Object{
				//Name: "zhenxinma",
			})
		}()
	}

	poll.Put(&Object{
		//Name: "zhenxinma",
	})
	obj := poll.Get().(*Object)
	// 回收回来的结构体还时一样, 所以这个会在什么情况下会有用到、
	fmt.Println(obj)
}

func main() {
	//testSyncPool()
	//fmt.Println(unsafe.Sizeof(Object{}))
	//fmt.Println(64 % 128)
	//obj := Obj{}
	//fmt.Println(unsafe.Sizeof(obj))
	//fmt.Println(uintptr(unsafe.Pointer(&obj.A)))
	//fmt.Println(uintptr(unsafe.Pointer(&obj.B)))
	// 824634564344
	// 824634564392
	// 4344+48

	//atomic.CompareAndSwapPointer()
	// 如果为空,那能否写入东西呢
	//writeAny(&Object{A: "zhenxinma"})

	// limit: 每秒可以向桶中放入多少token.
	// bocuket: 表示桶的大小.
	limiter := rate.NewLimiter(2, 3)
	for {
		limiter.WaitN(context.Background(), 1)
		fmt.Println(time.Now().Unix())
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	lock := sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()
}

func writeAny(a any) {
	array := make([]*Object, 8)
	va := array[4]
	*(*any)(unsafe.Pointer(va)) = a
	fmt.Printf("%+v\n", va)
}

type Obj struct {
	A [24]byte
	B [24]byte
}
