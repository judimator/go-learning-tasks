package two_sum

import (
	"reflect"
	"testing"
)

type TestData struct {
	nums   []int
	target int
	output []int
}

var tests = []TestData{
	{
		[]int{2, 7, 11, 15},
		9,
		[]int{0, 1},
	},
}

func TestTwoSum(t *testing.T) {
	for _, data := range tests {
		if result := twoSum(data.nums, data.target); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, twoSum")
		}
	}
}
