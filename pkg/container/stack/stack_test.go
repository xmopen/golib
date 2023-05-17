package stack

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	s := New(8)
	s.Push("hello word1")
	s.Push("hello word2")
	s.Push("hello word3")
	//s.Push("hello word4")
	//s.Push("hello word5")
	item1, _ := s.Pop(0)
	fmt.Println(item1)
	item2, _ := s.Pop(1)
	fmt.Println(item2)
	item3, _ := s.Pop(2)
	fmt.Println(item3)
}
