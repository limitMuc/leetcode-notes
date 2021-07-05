package leetcode

// 时间复杂度: O(1)
// 空间复杂度: O(1)
func deleteNode(node *ListNode) {
	if node == nil || node.Next == nil {
		return
	}

	node.Val = node.Next.Val
	delNode := node.Next
	node.Next = delNode.Next
	delNode.Next = nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
