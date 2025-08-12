package palindrominkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	arrLen := len(arr)
	for i := 0; i < arrLen/2; i++ {
		if arr[i] != arr[arrLen-1-i] {
			return false
		}
	}
	return true
}

func isPalindromePointers(head *ListNode) bool {
	// found middle
	slowPtr := head // will be middle
	fastPtr := head // will be end
	for slowPtr.Next != nil && fastPtr.Next != nil {
		slowPtr = slowPtr.Next
		if fastPtr.Next.Next == nil {
			break
		}
		fastPtr = fastPtr.Next.Next
	}

	// second reverse second part
	tmp := *slowPtr
	newCopy := tmp
	prev := &newCopy

	tmp = *slowPtr.Next
	newCopy = tmp
	cur := &newCopy

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	//compare two parts
	for fastPtr != nil && head != nil && (fastPtr.Next != head.Next || fastPtr.Next != head) {
		if fastPtr.Val != head.Val {
			return false
		}
		fastPtr = fastPtr.Next
		head = head.Next
	}

	return true
}
