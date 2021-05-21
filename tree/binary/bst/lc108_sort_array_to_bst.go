package bst

import (
	"leetcode/tree/binary/bt"
)

// 给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。

// 高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。

// https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/

func sortedArrayToBST(nums []int) *bt.TreeNode {
	return sort(nums, 0 , len(nums) - 1)
}

func sort(nums []int, left, right int) *bt.TreeNode {

	if left > right {
		return nil
	}
	mid := (left + right + 1)/2
	root := &bt.TreeNode{Val: nums[mid]}
	root.Left = sort(nums, left, mid - 1)
	root.Right = sort(nums, mid+1, right)
	return root
}




