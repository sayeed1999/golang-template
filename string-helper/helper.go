package stringhelper

import (
	"sort"
	"strings"
)

func SortString(s string) string {
	letters := strings.Split(s, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}

func GetLowercaseCharFrequencyMap(s string) [26]int {
	lenx := len(s)
	frequencyMap := [26]int{}

	for i := 0; i < lenx; i++ {
		if s[i] >= 97 || s[i] < 122 {
			frequencyMap[s[i]-97]++
		}
	}
	return frequencyMap
}

func GetUppercaseCharFrequencyMap(s string) [26]int {
	lenx := len(s)
	frequencyMap := [26]int{}

	for i := 0; i < lenx; i++ {
		if s[i] >= 65 || s[i] < 90 {
			frequencyMap[s[i]-65]++
		}
	}
	return frequencyMap
}
