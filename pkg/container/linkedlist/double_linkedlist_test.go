package linkedlist

import (
	"fmt"
	"runtime"
	"testing"
)

func TestPushHeadAndTail(t *testing.T) {
	list := NewDoubleLinkedList()
	list.PushWithHead(2)
	list.PushWithHead(1)
	list.PushWithHead(0)
	// 0 1 2
	printDoubleLinkedListWithHead(list)
	runtime.GC()
	// 0 1 2
	list = NewDoubleLinkedList()
	list.PushWithTail(0)
	list.PushWithTail(1)
	list.PushWithTail(2)
	printDoubleLinkedListWithTail(list)
}

func printDoubleLinkedListWithHead(list IDoubleLinkedList) {
	for {
		head := list.RemoveFromHead()
		if head == nil {
			break
		}
		fmt.Println(head)
	}
	fmt.Printf("size:%v\n", list.Size())
}

func printDoubleLinkedListWithTail(list IDoubleLinkedList) {
	for {
		head := list.RemoveFromTail()
		if head == nil {
			break
		}
		fmt.Println(head)
	}
	fmt.Printf("size:%v\n", list.Size())
}

func TestDoubleLinkedListPeek(t *testing.T) {
	list := NewDoubleLinkedList()
	list.PushWithHead(2)
	list.PushWithHead(1)
	list.PushWithHead(0)
	fmt.Println(list.PeekFromHead())
	fmt.Println(list.PeekFromTail())
}
