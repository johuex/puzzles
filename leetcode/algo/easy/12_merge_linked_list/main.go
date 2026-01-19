package mergelinkedlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// no need to merge if one of list is empty
	if list1 == nil && list2 != nil {
		return list2
	}
	if list1 != nil && list2 == nil {
		return list1
	}
	if list1 == nil && list2 == nil {
		return nil
	}

	mergeHead := &ListNode{}

	leftActualItem := list1
	rightActualItem := list2
	// initial of result list
	if leftActualItem.Val < rightActualItem.Val {
		mergeHead.Val = leftActualItem.Val
		leftActualItem = leftActualItem.Next
	} else {
		mergeHead.Val = rightActualItem.Val
		rightActualItem = rightActualItem.Next
	}

	mActualItem := mergeHead
	for {
		nextMergeItem := &ListNode{}
		mActualItem.Next = nextMergeItem
		mActualItem = nextMergeItem
		if rightActualItem == nil || (leftActualItem != nil && leftActualItem.Val < rightActualItem.Val) {
			mActualItem.Val = leftActualItem.Val
			leftActualItem = leftActualItem.Next
		} else {
			mActualItem.Val = rightActualItem.Val
			rightActualItem = rightActualItem.Next
		}
		if leftActualItem == nil && rightActualItem == nil {
			break
		}
	}
	return mergeHead
}
