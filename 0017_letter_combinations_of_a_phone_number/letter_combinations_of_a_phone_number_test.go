package letter_combinations_of_a_phone_number

import (
	"reflect"
	"testing"
)

type TestData struct {
	input  string
	output []string
}

var tests = []TestData{
	{
		"23",
		[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
	},
}

func TestThreeSumClosest(t *testing.T) {
	for _, data := range tests {
		if result := letterCombinations(data.input); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, letterCombinations")
		}
	}
}
