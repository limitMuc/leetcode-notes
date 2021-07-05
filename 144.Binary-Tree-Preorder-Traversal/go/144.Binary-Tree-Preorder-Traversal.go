package leetcode

// 二叉树的前序遍历
// 时间复杂度: O(n), n为树的节点个数
// 空间复杂度: O(h), h为树的高度
func preorderTraversal(root *TreeNode) (vals []int) {
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
