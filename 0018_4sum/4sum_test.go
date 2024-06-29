package foursum

import (
	"reflect"
	"testing"
)

type TestData struct {
	input  []int
	target int
	output [][]int
}

var tests = []TestData{
	{
		[]int{1, 0, -1, 0, -2, 2},
		0,
		[][]int{{-1, -1, 2}, {-1, 0, 1}},
	},
	{
		[]int{2, 2, 2, 2, 2},
		8,
		[][]int{{2, 2, 2, 2}},
	},
}

func TestFourSum(t *testing.T) {
	for _, data := range tests {
		if result := fourSum(data.input, data.target); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, fourSum")
		}
	}
}
