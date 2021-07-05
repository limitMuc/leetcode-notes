package leetcode

import (
	"fmt"
	"strconv"
)

// 时间复杂度: O(n), n为树中的节点个数
// 空间复杂度: O(h), h为树的高度
func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	if root == nil {
		return res
	}

	if root.Left == nil && root.Right == nil {
		res = append(res, strconv.Itoa(root.Val))
	}

	leftS := binaryTreePaths(root.Left)
	for _, str := range leftS {
		res = append(res, fmt.Sprintf("%d->%s", root.Val, str))
	}

	rightS := binaryTreePaths(root.Right)
	for _, str := range rightS {
		res = append(res, fmt.Sprintf("%d->%s", root.Val, str))
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
