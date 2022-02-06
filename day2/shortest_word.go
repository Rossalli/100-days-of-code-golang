package main

import (
	"strings"
)

func ShortestWord(str string) int {
	//split string by withspace
	s := strings.Fields(str)

	//iterate slice and verify short
	short := 999
	for _, w := range s {
		if len(w) < short {
			short = len(w)
		}
	}
	return short
}
