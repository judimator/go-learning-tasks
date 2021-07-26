package minimum_path_sum

import (
	"testing"
)

type TestData struct {
	input  [][]int
	output int
}

var tests = []TestData{
	{
		[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
		7,
	},
	{
		[][]int{{1, 2, 3}, {4, 5, 6}},
		12,
	},
	{
		[][]int{{1, 2}, {5, 6}, {1, 1}},
		8,
	},
}

func TestMinPathSum(t *testing.T) {
	for _, data := range tests {
		if result := minPathSum(data.input); result != data.output {
			t.Error("Failed, TestConvert")
		}
	}
}
