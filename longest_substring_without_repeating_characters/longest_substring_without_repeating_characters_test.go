package longest_substring_without_repeating_characters

import (
	"reflect"
	"testing"
)

type TestData struct {
	input  string
	output int
}

var tests = []TestData{
	{
		"abcabcbb",
		3,
	},
	{
		"bbbbb",
		1,
	},
	{
		"pwwkew",
		3,
	},
	{
		"",
		0,
	},
}

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	for _, data := range tests {
		if result := lengthOfLongestSubstring(data.input); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, lengthOfLongestSubstring")
		}
	}
}
