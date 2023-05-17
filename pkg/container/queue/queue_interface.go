// Package queue
// Create  2023-03-13 19:06:29
package queue

// Queue 队列接口.
type Queue interface {
	Push(val any)
	Pop() (any, error)
	Len() int64
	String() string
}
