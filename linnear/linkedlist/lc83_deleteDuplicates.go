package linkedlist


// 存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。
// 返回同样按升序排列的结果链表。

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/

// 1 3 3 4 6 6 -> 1 3 4 6
func deleteDuplicates(head *ListNode) *ListNode {

	pri := head
	for head != nil && head.Next != nil  {

		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}

	return pri
}

// 1 3 3 4 6 6 -> 1 3 4 6
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	 head.Next = deleteDuplicates(head.Next)
	 if head.Val == head.Next.Val {
	 	head = head.Next
	 }
	 return head
}
