package zigzag_conversion

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var output string
	var counter = 0
	var reverse = false
	var data = make([][]string, numRows)

	for _, value := range s {
		data[counter] = append(data[counter], string(value))

		if reverse == true {
			counter--
		} else {
			counter++
		}

		if counter == numRows-1 || counter == 0 {
			reverse = !reverse
		}
	}

	for _, parts := range data {
		output += strings.Join(parts, "")
	}
	return output
}
