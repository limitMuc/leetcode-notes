package leetcode

// 时间复杂度: O(n)
// 空间复杂度: O(1)
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Next: head,
	}
	for pre := dummyHead; pre.Next != nil && pre.Next.Next != nil; {
		lnode := pre.Next
		rnode := lnode.Next
		next := rnode.Next

		pre.Next = rnode
		rnode.Next = lnode
		lnode.Next = next

		pre = lnode // 交换完成后，pre就应该位于下个要交换的一对节点的前置节点
	}

	retNode := dummyHead.Next
	dummyHead.Next = nil

	return retNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}
