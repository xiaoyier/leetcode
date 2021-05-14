package linkedlist

// 给定一个头结点为 head 的非空单链表，返回链表的中间结点。
// 如果有两个中间结点，则返回第二个中间结点。

// https://leetcode-cn.com/problems/middle-of-the-linked-list/

// 1 2 3 4 5 -> 3 4 5
// 1 2 3 4 5 6  -> 4 5 6
func middleNode(head *ListNode) *ListNode {

	var nodes = make([]*ListNode, 0)
	for head != nil {

		nodes = append(nodes, head)
		head = head.Next
	}

	var index int = len(nodes)/2
	return nodes[index]
}

func middleNode2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	p := head
	q := head
	for q != nil && q.Next != nil {
		q = q.Next.Next
		p = p.Next
	}

	return p
	//for q != nil {
	//	q = q.Next
	//	if q != nil {
	//		q = q.Next
	//		p = p.Next
	//	}
	//}
	//return p
}

