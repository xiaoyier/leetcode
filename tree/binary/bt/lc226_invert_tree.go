package bt

// 翻转一棵二叉树。

// https://leetcode-cn.com/problems/invert-binary-tree/

func invertTree(root *TreeNode) *TreeNode {
	return invertTree4(root)
}

func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}


func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left

	return root
}


func invertTree3(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	invertTree(root.Left)
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)

	return root
}


func invertTree4(root *TreeNode) *TreeNode {

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		node.Left, node.Right = node.Right, node.Left

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}
