package median_of_two_sorted_arrays

import (
	"testing"
)

type TestData struct {
	array1 []int
	array2 []int
	output float64
}

var tests = []TestData{
	{
		[]int{1, 3},
		[]int{2},
		2,
	},
	{
		[]int{1, 2, 3, 4, 5, 6},
		[]int{1, 2, 3},
		3.0,
	},
	{
		[]int{0, 0},
		[]int{0, 0},
		0.0,
	},
	{
		[]int{},
		[]int{1},
		1.0,
	},
}

func TestFindMedianSortedArrays(t *testing.T) {
	for _, data := range tests {
		if result := findMedianSortedArrays(data.array1, data.array2); result != data.output {
			t.Error("Failed, findMedianSortedArrays")
		}
	}
}
