package add_two_sum_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoSumNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	val1, val2, carry, current := 0, 0, 0, result
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			val1 = 0
		} else {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			val2 = 0
		} else {
			val2 = l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{Val: (val1 + val2 + carry) % 10}
		current = current.Next
		carry = (val1 + val2 + carry) / 10
	}
	return result.Next
}
