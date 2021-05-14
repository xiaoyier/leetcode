package linkedlist

import (
	"testing"
)

func TestDeleteNode(t *testing.T) {
	node := &ListNode{
		Val: 5,
	}

	for i := 4; i > 0; i-- {
		next := &ListNode{
			Val: i,
		}
		node.Next = next
		node = next
	}

	deleteNode(node)
}
