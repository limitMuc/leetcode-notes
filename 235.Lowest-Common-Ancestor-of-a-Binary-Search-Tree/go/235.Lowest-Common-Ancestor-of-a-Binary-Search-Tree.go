package leetcode

// 时间复杂度: O(lgn), 其中n为树的节点个数
// 空间复杂度: O(h), 其中h为树的高度
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
