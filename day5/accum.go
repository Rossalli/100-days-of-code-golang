package main

import "strings"

func Accum(s string) string {
	accum := ""
	arr := []rune(s)

	for i := 0; i < len(arr); i++ {

		letter := strings.ToLower(string(arr[i]))
		accum = accum + strings.ToUpper(letter) + strings.Repeat(letter, i) + "-"
	}
	return accum[:len(accum)-1]
}

func AccumRefactored(s string) string {
	arr := strings.Split(s, "")
	for i, letter := range arr {
		arr[i] = strings.Title(strings.Repeat(strings.ToLower(letter), i+1))
	}
	return strings.Join(arr, "-")
}
