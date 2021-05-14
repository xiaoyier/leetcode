package linkedlist

import (
	"fmt"
	"testing"
)

func TestHasCycle(t *testing.T) {

	pri := &ListNode{
		Val: 5,
	}

	head := pri
	fmt.Printf("%p, %p\n", &pri, head)
	nodes := make([]*ListNode, 4)
	for i := 4; i > 0; i-- {
		node := &ListNode{
			Val: i,
		}
		head.Next = node
		head = node
		nodes[i-1] = node
	}
	nodes[1].Next = nodes[3]
	fmt.Printf("%p, %p\n", &pri, head)
	fmt.Println(hasCycle(pri))
	*head = ListNode{
		Val: 100,
	}
	fmt.Printf("%p, %p\n", &pri, head)
	fmt.Println(pri, head)

	a := 5
	b := &a
	c := b
	a = 10
	*b = 15
	fmt.Println(a, *b, *c)
	fmt.Printf("%p, %p, %p\n", &a, b, c)

	//x := ListNode{
	//	Val: 111,
	//}
	y := &ListNode{
		Val: 111,
	}
	z := y
	//x = ListNode{
	//	Val:222,
	//}
	z = &ListNode{
		Val: 333,
	}
	fmt.Println(*y,*z)
	fmt.Printf("%p, %p", y, z)

}

