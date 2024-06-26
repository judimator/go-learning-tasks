package letter_combinations_of_a_phone_number

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var result []string
	lookup := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	getCombinations(0, []rune(digits), "", lookup, &result)

	return result
}

func getCombinations(i int, digits []rune, memo string, lookup map[rune]string, result *[]string) {
	if i >= len(digits) {
		*result = append(*result, memo)
		return
	}

	alphabet := lookup[digits[i]]

	for _, char := range alphabet {
		tmp := memo + string(char)
		getCombinations(i+1, digits, tmp, lookup, result)
	}
}
