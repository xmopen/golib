package linkedlist

import (
	"fmt"
	"testing"
)

func TestLinkedListAdd(t *testing.T) {
	list := New()
	for i := 0; i < 5; i++ {
		list.Add(i)
	}
	fmt.Println("print list struct")
	for {
		next := list.Next()
		if next == nil {
			return
		}
		fmt.Println(next)
		list.Remove()
	}
}

func TestLinkedListRemove(t *testing.T) {
	list := New()
	for i := 0; i < 5; i++ {
		list.Add(i)
	}
	fmt.Println("print list struct")
	list.Remove()
	list.Remove()
	list.Remove()
	// 3 4
	for {
		next := list.Next()
		if next == nil {
			return
		}
		fmt.Println(next)
		list.Remove()
	}
}

func TestLinkedListSizeAndEmpty(t *testing.T) {
	list := New()
	for i := 0; i < 5; i++ {
		list.Add(i)
	}
	list.Remove()
	fmt.Println(list.Empty())
	fmt.Println(list.Size())
	list.Remove()
	list.Remove()
	list.Remove()
	list.Remove()
	list.Remove()
	fmt.Println(list.Empty())
}
