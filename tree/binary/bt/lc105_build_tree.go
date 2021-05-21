package bt


// 根据一棵树的前序遍历与中序遍历构造二叉树。

// 注意: 你可以假设树中没有重复的元素。

// 前序遍历 preorder = [3,9,20,15,7]
// 中序遍历 inorder = [9,3,15,20,7]
func BuildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	mid := preorder[0]
	preorder = preorder[1:]
	index := -1
	for i, v := range inorder {
		if v == mid {
			index = i
			break
		}
	}

	root := &TreeNode{Val: mid}
	if index > 0 {
		root.Left = BuildTree1(preorder[:index], inorder[:index])
	}

	if len(preorder) > 0 {
		root.Right = BuildTree1(preorder[index:], inorder[index+1:])
	}
	return root
}
