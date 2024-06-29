package remove_nth_node_from_end_of_list

import (
	"reflect"
	"testing"
)

type TestData struct {
	input  *ListNode
	n      int
	output *ListNode
}

func TestRemoveNthFromEnd1(t *testing.T) {
	node5 := ListNode{5, nil}
	node4 := ListNode{4, &node5}
	node3 := ListNode{3, &node4}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}

	node55 := ListNode{5, nil}
	node33 := ListNode{3, &node55}
	node22 := ListNode{2, &node33}
	node11 := ListNode{1, &node22}

	var tests = []TestData{
		{
			&node1,
			2,
			&node11,
		},
	}

	for _, data := range tests {
		if result := removeNthFromEnd(data.input, data.n); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, removeNthFromEnd1")
		}
	}
}

func TestRemoveNthFromEnd2(t *testing.T) {
	node2 := ListNode{2, nil}
	node1 := ListNode{1, &node2}

	node11 := ListNode{1, nil}

	var tests = []TestData{
		{
			&node1,
			1,
			&node11,
		},
	}

	for _, data := range tests {
		if result := removeNthFromEnd(data.input, data.n); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, removeNthFromEnd2")
		}
	}
}

func TestRemoveNthFromEnd3(t *testing.T) {
	node1 := ListNode{1, nil}

	var tests = []TestData{
		{
			&node1,
			1,
			nil,
		},
	}

	for _, data := range tests {
		if result := removeNthFromEnd(data.input, data.n); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, removeNthFromEnd3")
		}
	}
}

func TestRemoveNthFromEnd4(t *testing.T) {
	node2 := ListNode{2, nil}
	node1 := ListNode{1, &node2}

	node22 := ListNode{2, nil}

	var tests = []TestData{
		{
			&node1,
			2,
			&node22,
		},
	}

	for _, data := range tests {
		if result := removeNthFromEnd(data.input, data.n); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, removeNthFromEnd4")
		}
	}
}

func TestRemoveNthFromEnd5(t *testing.T) {
	node3 := ListNode{3, nil}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}

	node22 := ListNode{2, nil}
	node11 := ListNode{1, &node22}

	var tests = []TestData{
		{
			&node1,
			1,
			&node11,
		},
	}

	for _, data := range tests {
		if result := removeNthFromEnd(data.input, data.n); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, removeNthFromEnd5")
		}
	}
}
