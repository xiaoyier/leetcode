package bt


// 给你二叉树的根结点 root ，请你将它展开为一个单链表：
//
//展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
//展开后的单链表应该与二叉树 先序遍历 顺序相同。

// https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/

func Flatten(root *TreeNode)  {
	if root == nil {
		return
	}

	stack := []*TreeNode{}
	var preNode *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {

			if preNode != nil {
				preNode.Right = root
				preNode.Left = nil
			}
			if root.Right != nil {
				stack = append(stack, root.Right)
			}

			preNode = root
			root = root.Left
		}

		if len(stack) == 0 {
			return
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
}