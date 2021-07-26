package zigzag_conversion

import (
	"testing"
)

type TestData struct {
	input   string
	numRows int
	output  string
}

var tests = []TestData{
	{
		"PAYPALISHIRING",
		3,
		"PAHNAPLSIIGYIR",
	},
	{
		"PAYPALISHIRING",
		4,
		"PINALSIGYAHRPI",
	},
	{
		"A",
		1,
		"A",
	},
}

func TestConvert(t *testing.T) {
	for _, data := range tests {
		if result := convert(data.input, data.numRows); result != data.output {
			t.Error("Failed, TestConvert")
		}
	}
}
