package leetcode

// 二叉树的后序遍历
// 时间复杂度: O(n), n为树的节点个数
// 空间复杂度: O(h), h为树的高度
func postorderTraversal(root *TreeNode) (res []int) {
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)
	return
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}