package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Peek())
	fmt.Println(stack.Pop())
	stack.Push(4)
	ele := stack.Pop()
	fmt.Println(ele)
}
