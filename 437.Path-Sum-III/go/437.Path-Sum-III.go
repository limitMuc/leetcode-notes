package leetcode

// 时间复杂度: O(n), n为树的节点个数
// 空间复杂度: O(h), h为树的高度
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return findPath(root, targetSum) +
		pathSum(root.Left, targetSum) +
		pathSum(root.Right, targetSum)
}

// 在以node为根节点的二叉树中,寻找包含node的路径,和为sum
// 返回这样的路径个数
func findPath(node *TreeNode, targetSum int) int {
	if node == nil {
		return 0
	}

	num := 0
	if targetSum == node.Val {
		num++
	}

	num += findPath(node.Left, targetSum-node.Val)
	num += findPath(node.Right, targetSum-node.Val)
	return num
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
