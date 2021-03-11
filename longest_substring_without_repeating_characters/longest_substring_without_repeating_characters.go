package longest_substring_without_repeating_characters

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	tmp := make(map[string]int)
	longest := 0
	chars := strings.Split(s, "")
	copyChairs := make([]string, len(chars))

	copy(copyChairs, chars)

	for range chars {
		for _, charInner := range copyChairs {
			_, ok := tmp[charInner]
			if ok == false {
				tmp[charInner] = 1
				tmpLongest := len(tmp)

				if tmpLongest > longest {
					longest = tmpLongest
				}
			} else {
				copyChairs = remove(copyChairs, 0)
				tmp = map[string]int{}

				break
			}
		}
	}
	return longest
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
