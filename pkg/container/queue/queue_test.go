package queue

import (
	"fmt"
	"testing"
)

// TestQueue 测试FIFO特性
func TestQueue(t *testing.T) {
	queue := New()
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)
	fmt.Printf("size:[%+v]\n", queue.Size())
	fmt.Println(queue.Poll())
	fmt.Println(queue.Poll())
	fmt.Println(queue.Poll())
	fmt.Println(queue.Poll())
}
