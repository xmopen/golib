// Package queue
// Create  2023-03-13 19:06:29
package queue

import "github.com/xmopen/golib/pkg/container"

// Queue 队列接口.
// TODO: 后续废弃掉
//type Queue interface {
//	Push(val any)
//	Pop() (any, error)
//	Len() int64
//	String() string
//}

// IQueue Queue interface
type IQueue interface {
	container.IContainer
	Poll() any
	Peek() any
}
