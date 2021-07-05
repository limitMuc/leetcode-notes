package leetcode

// 时间复杂度: O(n), n为树中节点个数
// 空间复杂度: O(h), h为树的高度
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
