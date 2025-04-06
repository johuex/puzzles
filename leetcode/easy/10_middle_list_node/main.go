package middlelistnode

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	nodesArr := make([]*ListNode, 0)
	for head != nil {
		nodesArr = append(nodesArr, head)
		head = head.Next
	}
	if len(nodesArr) == 1 {
		return nodesArr[0]
	}

	return nodesArr[len(nodesArr)/2]
}

func middleNodePointers(head *ListNode) *ListNode {
	slowPtr := head
	fastPtr := head
	for slowPtr.Next != nil && fastPtr.Next != nil {
		slowPtr = slowPtr.Next
		if fastPtr.Next.Next == nil {
			break
		}
		fastPtr = fastPtr.Next.Next
	}
	return slowPtr
}
