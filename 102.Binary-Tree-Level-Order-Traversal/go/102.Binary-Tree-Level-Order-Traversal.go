package leetcode

// 二叉树的层序遍历
// 时间复杂度: O(n), n为树的节点个数
// 空间复杂度: O(n)
func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	levelNodes := []*TreeNode{root}
	for i := 0; len(levelNodes) > 0; i++ {
		nextLevelNodes := []*TreeNode{}
		ret = append(ret, []int{})
		for j := 0; j < len(levelNodes); j++ {
			node := levelNodes[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				nextLevelNodes = append(nextLevelNodes, node.Left)
			}
			if node.Right != nil {
				nextLevelNodes = append(nextLevelNodes, node.Right)
			}
		}
		levelNodes = nextLevelNodes
	}
	return ret
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}