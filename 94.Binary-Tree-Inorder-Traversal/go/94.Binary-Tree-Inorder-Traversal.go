package leetcode

// 二叉树的中序遍历
// 时间复杂度: O(n), n为树的节点个数
// 空间复杂度: O(h), h为树的高度
func inorderTraversal(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}