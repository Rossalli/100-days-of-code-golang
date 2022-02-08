package main

import (
	"sort"
	"strings"
)

func TwoSort(arr []string) string {

	word := ""

	sort.Slice(arr, func(i int, j int) bool {
		return arr[i] < arr[j]
	})

	s := []rune(arr[0])

	for i := 0; i < len(s); i++ {
		word = word + string(s[i]) + "***"
	}

	return word[:len(word)-3]

}

func TwoSortRefactored(arr []string) string {
	sort.Strings(arr)
	chars := strings.Split(arr[0], "")
	return strings.Join(chars, "***")
}
