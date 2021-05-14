package linkedlist


// 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

// https://leetcode-cn.com/problems/remove-linked-list-elements/

func removeElements1(head *ListNode, val int) *ListNode {

	newHead := &ListNode{
		Val: 0,
		Next: head,
	}
	curr := newHead
	for head != nil {

		if head.Val == val {
			curr.Next = head.Next
		} else {
			curr = head
		}

		head = head.Next
	}
	return newHead.Next
}

// 如何处理最后一个节点的删除
func removeElements2(head *ListNode, val int) *ListNode {
	var newHead *ListNode = nil
	for head != nil {
		if head.Val == val {
			if head.Next != nil {
				head.Val = head.Next.Val
				head.Next = head.Next.Next
				head = head.Next.Next
			} else {
				head = nil //???
			}
		} else {
			if newHead == nil {
				newHead = head
			}
			head = head.Next
		}
	}
	return newHead
}

func removeElements3(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	head.Next = removeElements3(head.Next, val)
	if head.Val == val {
		return head.Next
	} else {
		return head
	}
}
