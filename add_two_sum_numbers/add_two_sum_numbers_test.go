package add_two_sum_numbers

import (
	"reflect"
	"testing"
)

type TestData struct {
	l1     *ListNode
	l2     *ListNode
	output *ListNode
}

var tests = []TestData{
	{
		castToListNode([]int{2, 4, 3}),
		castToListNode([]int{5, 6, 4}),
		castToListNode([]int{7, 0, 8}),
	},
	{
		castToListNode([]int{0}),
		castToListNode([]int{0}),
		castToListNode([]int{0}),
	},
	{
		castToListNode([]int{9, 9, 9, 9, 9, 9, 9}),
		castToListNode([]int{9, 9, 9, 9}),
		castToListNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
	},
}

func castToListNode(array []int) *ListNode {
	list := &ListNode{}
	tmp := list
	for _, v := range array {
		tmp.Next = &ListNode{Val: v}
		tmp = tmp.Next
	}
	return list.Next
}

func TestAddTwoSumNumbers(t *testing.T) {
	for _, data := range tests {
		if result := addTwoSumNumbers(data.l1, data.l2); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, addTwoSumNumbers")
		}
	}
}
