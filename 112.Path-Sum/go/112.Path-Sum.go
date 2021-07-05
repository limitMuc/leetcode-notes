package leetcode

// 时间复杂度: O(n), n为树的节点个数
// 空间复杂度: O(h), h为树的高度
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if (root.Left == nil && root.Right == nil) {
		return root.Val == targetSum
	}
	return hasPathSum(root.Left, targetSum - root.Val) ||
		hasPathSum(root.Right, targetSum - root.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
