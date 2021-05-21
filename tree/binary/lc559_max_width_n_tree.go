package binary

// 给定一个 N 叉树，找到其最大深度。
//
//最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。
//
//N 叉树输入按层序遍历序列化表示，每组子节点由空值分隔（请参见示例）。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree/

func maxDepthN(root *Node) int {

	if root == nil {
		return 0
	}
	q := []*Node{root}
	depth, size := 0, 1
	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		for _, item := range node.Children {
			q = append(q, item)
		}
		size--
		if size == 0 {
			depth++
			size =len(q)
		}
	}
	return depth
}