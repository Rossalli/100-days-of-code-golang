package main

import (
	"sort"
)

func Comp(arr1 []int, arr2 []int) bool {

	if arr1 == nil || arr2 == nil || len(arr1) != len(arr2) {
		return false
	}

	for i, n := range arr1 {
		arr1[i] = n * n
	}

	sort.Ints(arr1)
	sort.Ints(arr2)

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
