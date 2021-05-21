package n

// 给定一个 N 叉树，返回其节点值的 后序遍历 。
//
//N 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。

// https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/


func postorder(root *Node) []int {
	vals := []int{}
	stack := []*Node{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			if len(root.Children) > 0 {
				tmp := root.Children
				root.Children = tmp[1:]
				root = tmp[0]
			} else {
				root = nil
			}
		}

		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if len(root.Children) == 0 {
			vals = append(vals, root.Val)
			root = nil
		}
	}
	return vals
}
func postorder1(root *Node) []int {
	vals := []int{}
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		for _, item := range node.Children {
			dfs(item)
		}
		vals = append(vals, node.Val)
	}
	dfs(root)
	return vals
}
