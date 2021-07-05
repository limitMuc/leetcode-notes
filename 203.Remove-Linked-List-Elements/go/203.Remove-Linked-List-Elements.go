package leetcode

// 设立链表的虚拟头结点
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{
		Next: head,
	}

	for cur := dummyHead; cur.Next != nil; {
		if cur.Next.Val == val {
			delNode := cur.Next
			cur.Next = delNode.Next
			delNode.Next = nil // 相当于删除节点，触发gc
		} else {
			cur = cur.Next
		}
	}

	retNode := dummyHead.Next
	dummyHead.Next = nil
	return retNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}
