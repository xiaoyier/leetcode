package binary

import (
	"fmt"
	"testing"
)

func TestInverTree(t *testing.T) {
	node4 := &TreeNode{Val: 4}
	node2 := &TreeNode{Val: 2}
	node7 := &TreeNode{Val: 7}
	node1 := &TreeNode{Val: 1}
	node3 := &TreeNode{Val: 3}
	node6 := &TreeNode{Val: 6}
	node9 := &TreeNode{Val: 9}
	node4.Left = node2
	node4.Right = node7
	node2.Left = node1
	node2.Right = node3
	node7.Left = node6
	node7.Right = node9
	invertTree(node4)

	FormatTree(node4)
}

func TestDeleteNode(t *testing.T) {

	//node5 := &TreeNode{Val: 5}
	node3 := &TreeNode{Val: 3}
	//node6 := &TreeNode{Val: 6}
	node2 := &TreeNode{Val: 2}
	node4 := &TreeNode{Val: 4}
	node1 := &TreeNode{Val: 1}
	//node7 := &TreeNode{Val: 7}
	//node5.Left = node3
	//node5.Right = node6
	node3.Left = node2
	node3.Right = node4
	node2.Left = node1
	//node6.Left = node0
	//node6.Right = node7

	FormatTree(node3)
	fmt.Println()

	node := deleteNode(node3, 1)
	FormatTree(node)
}


func FormatTree(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	FormatTree(root.Left)
	FormatTree(root.Right)
}

func TestPostOrder(t *testing.T) {

	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}
	node3 := &Node{Val: 3, Children: []*Node{node5,node6}}
	node2 := &Node{Val: 2}
	node4 := &Node{Val: 4}
	node1 := &Node{Val: 1, Children: []*Node{node3,node2,node4}}
	result := postorder(node1)
	fmt.Println(result)
}



