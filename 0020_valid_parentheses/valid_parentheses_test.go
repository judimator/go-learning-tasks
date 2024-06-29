package valid_parentheses

import "testing"

type TestData struct {
	input  string
	output bool
}

var tests = []TestData{
	//{
	//	"()",
	//	true,
	//},
	//{
	//	"()[]{}",
	//	true,
	//},
	//{
	//	"(]",
	//	false,
	//},
	//{
	//	"[",
	//	false,
	//},
	{
		"(){}}{",
		false,
	},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, data := range tests {
		if result := isValid(data.input); result != data.output {
			t.Error("Failed, isValid")
		}
	}
}
