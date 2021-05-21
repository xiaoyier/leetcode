package bt


// 根据一棵树的中序遍历与后序遍历构造二叉树。

// 注意: 你可以假设树中没有重复的元素。

// https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/


// 中序遍历 inorder = [9,3,15,20,7]
// 后序遍历 postorder = [9,15,7,20,3]
func BuildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	mid := postorder[len(postorder)-1]
	postorder = postorder[:len(postorder)-1]
	index := -1
	for i, v := range inorder {
		if v == mid{
			index = i
			break
		}
	}
	root := &TreeNode{Val: mid}
	if index > 0 {
		root.Left = BuildTree(inorder[:index], postorder[:index])
	}
	root.Right = BuildTree(inorder[index+1:], postorder[index:])

	return root
}