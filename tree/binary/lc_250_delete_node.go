package binary


func deleteNode(root *TreeNode, key int) *TreeNode {

	deleteNode, parent := findNode(root, key)
	if deleteNode == nil {
		return root
	}
	// 双度节点
	if deleteNode.Left != nil && deleteNode.Right != nil {
		successorNode, p := successor(deleteNode)
		deleteNode.Val = successorNode.Val
		deleteNode = successorNode
		parent = p
	}

	child := deleteNode.Left
	if deleteNode.Right != nil {
		child = deleteNode.Right
	}


	if deleteNode == root {
		root = child
	} else if deleteNode == parent.Left {
		parent.Left = child
	} else  {
		parent.Right = child
	}

	return root
}

// O(logn) 即 O(h)
func findNode(root *TreeNode, key int) (node *TreeNode, parent *TreeNode) {
	node = root
	for node != nil {
		if node.Val == key {
			return node, parent
		} else if node.Val > key {
			parent = node
			node = node.Left
		} else {
			parent = node
			node = node.Right
		}
	}
	return node, parent
}


func successor(node *TreeNode) (*TreeNode, *TreeNode) {
	if node == nil || node.Right == nil {
		return nil, nil
	}

	parent := node
	node = node.Right
	for node.Left != nil {
		parent = node
		node = node.Left
	}
	return node, parent
}