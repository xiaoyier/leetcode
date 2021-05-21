package binary

import (
	"fmt"
	"leetcode/tree/binary/bt"
	"testing"
)

func TestInverTree(t *testing.T) {
	node4 := &bt.TreeNode{Val: 4}
	node2 := &bt.TreeNode{Val: 2}
	node7 := &bt.TreeNode{Val: 7}
	node1 := &bt.TreeNode{Val: 1}
	node3 := &bt.TreeNode{Val: 3}
	node6 := &bt.TreeNode{Val: 6}
	node9 := &bt.TreeNode{Val: 9}
	node4.Left = node2
	node4.Right = node7
	node2.Left = node1
	node2.Right = node3
	node7.Left = node6
	node7.Right = node9
	//bt.invertTree(node4)

	FormatTree(node4)
}

func TestDeleteNode(t *testing.T) {

	//node5 := &TreeNode{Val: 5}
	node3 := &bt.TreeNode{Val: 3}
	//node6 := &TreeNode{Val: 6}
	node2 := &bt.TreeNode{Val: 2}
	node4 := &bt.TreeNode{Val: 4}
	node1 := &bt.TreeNode{Val: 1}
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

	//node := bst.deleteNode(node3, 1)
	//FormatTree(node)
}


func FormatTree(root *bt.TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	FormatTree(root.Left)
	FormatTree(root.Right)
}

func TestPostOrder(t *testing.T) {

	//node5 := &n.Node{Val: 5}
	//node6 := &n.Node{Val: 6}
	//node3 := &n.Node{Val: 3, Children: []*n.Node{node5,node6}}
	//node2 := &n.Node{Val: 2}
	//node4 := &n.Node{Val: 4}
	//node1 := &n.Node{Val: 1, Children: []*n.Node{node3,node2,node4}}
	//result := n.postorder(node1)
	//fmt.Println(result)
}


func TestFlatten(t *testing.T) {

	node1 := &bt.TreeNode{Val: 1}
	node2 := &bt.TreeNode{Val: 2}
	node3 := &bt.TreeNode{Val: 3}
	node4 := &bt.TreeNode{Val: 4}
	node5 := &bt.TreeNode{Val: 5}
	node6 := &bt.TreeNode{Val: 6}
	node1.Left = node2
	node1.Right = node5
	node2.Left = node3
	node2.Right = node4
	node5.Right = node6

	bt.Flatten(node1)
	//bt.invertTree(node4)

	FormatTree(node1)
}

func TestBuildTree(t *testing.T) {
	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}
	root := bt.BuildTree1(preorder,inorder)
	FormatTree(root)
}


