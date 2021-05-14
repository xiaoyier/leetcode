package linkedlist

// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

// https://leetcode-cn.com/problems/reverse-linked-list/

// 递归实现
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// 迭代实现
func reverseList2(head *ListNode) *ListNode {

	var preNode *ListNode = nil
	var tmpNode *ListNode = nil
	for head != nil {
		tmpNode = head.Next
		head.Next = preNode
		preNode = head
		head = tmpNode
	}
	return preNode
}
