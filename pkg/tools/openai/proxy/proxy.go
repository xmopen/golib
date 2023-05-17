// Package proxy
// Create  2023-03-28 22:50:37
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	// uint64高4位表示一级优先级 其余60位为二级优先级.
	var num uint64 = 0
	//high := uint64(num>>59) & uint64(num<<4-1) // 高4位.
	//low := uint64(num & uint64(num<<60-1))     // 低4位.
	//fmt.Println(high)
	//fmt.Println(low)
	// 比如现在一个任务优先级是10  入队时间是现在.1680076473
	// 第二个优先级是 20  入队时间是1680076474
	fmt.Println(time.Now().Unix())
	// 第一个任务入队.
	atomic.AddUint64(&num, uint64(10<<59))
	atomic.AddUint64(&num, uint64(1680076473))
	fmt.Printf("第一个元素scor:%+v\n", num)
	num = 0
	atomic.AddUint64(&num, uint64(10<<59))
	atomic.AddUint64(&num, uint64(1680076474))
	fmt.Printf("第一个元素scor:%+v\n", num)
	num = 0
	atomic.AddUint64(&num, uint64(20<<59))
	atomic.AddUint64(&num, uint64(1680076474))
	fmt.Printf("第一个元素scor:%+v\n", num)
}
