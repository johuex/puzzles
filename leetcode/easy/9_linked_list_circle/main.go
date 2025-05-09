package linkedlistcircle

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycleHash(head *ListNode) bool {
	pointers := make(map[*ListNode]struct{})
	pointers[head] = struct{}{}
	for head != nil {
		if _, ok := pointers[head.Next]; ok {
			return true
		} else {
			pointers[head] = struct{}{}
		}
		head = head.Next
	}
	return false
}

func hasCycleFastSlowPTR(head *ListNode) bool {
	slowPtr := head
	fastPtr := head
	for fastPtr != nil && fastPtr.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
		if fastPtr == slowPtr {
			return true
		}
	}
	return false
}
