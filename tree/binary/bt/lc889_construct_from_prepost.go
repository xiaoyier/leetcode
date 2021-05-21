package bt


// 返回与给定的前序和后序遍历匹配的任何二叉树。
// pre 和 post 遍历中的值是不同的正整数。

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/

// 输入：pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
// 输出：[1,2,3,4,5,6,7]

func constructFromPrePost(pre []int, post []int) *TreeNode {
	first := pre[0]
	root := &TreeNode{Val: first}
	if len(pre) == 1 {
		return root
	}

	pre = pre[1:]
	post = post[:len(post)-1]
	left := pre[0]
	right := post[len(post)-1]

	if left == right {
		//单子树
		root.Left = constructFromPrePost(pre, post)
	} else {
		preIndex, postIndex := -1, -1
		for i, v := range pre {
			if v == right {
				preIndex = i
				break
			}
		}
		for i, v := range post {
			if v == left {
				postIndex = i
				break
			}
		}

		//双子树
		root.Left = constructFromPrePost(pre[:preIndex], post[:postIndex+1])
		root.Right = constructFromPrePost(pre[preIndex:], post[postIndex+1:])
	}
	return root
}