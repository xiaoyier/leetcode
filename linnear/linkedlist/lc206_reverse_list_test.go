package linkedlist

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {

	head := &ListNode{
		Val: 5,
	}

	node := head
	for i := 4; i > 0; i-- {
		next := &ListNode{
			Val: i,
		}
		node.Next = next
		node = next
	}

	fmt.Println(reverseList2(head))
	fmt.Println(reverseList1(head))
}
