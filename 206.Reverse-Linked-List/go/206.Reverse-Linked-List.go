package leetcode

// 时间复杂度: O(n)
// 空间复杂度: O(1)
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for ; cur != nil; {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

type ListNode struct {
	Val  int
	Next *ListNode
}
