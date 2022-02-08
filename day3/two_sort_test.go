package main

import (
	"testing"
)

func TestTwoSort(t *testing.T) {

	t.Run("d***r***a***c***o dormiens nunquam titillandus", func(t *testing.T) {
		arr := []string{"draco", "dormiens", "nunquam", "titillandus"}
		result := TwoSort(arr)
		expected := "d***o***r***m***i***e***n***s"

		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})

	t.Run("you shall n***o***t pass", func(t *testing.T) {
		arr := []string{"you", "shall", "not", "pass"}
		result := TwoSort(arr)
		expected := "n***o***t"

		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})
}

func TestTwoSortRefactored(t *testing.T) {

	t.Run("d***r***a***c***o dormiens nunquam titillandus", func(t *testing.T) {
		arr := []string{"draco", "dormiens", "nunquam", "titillandus"}
		result := TwoSortRefactored(arr)
		expected := "d***o***r***m***i***e***n***s"

		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})

	t.Run("you shall n***o***t pass", func(t *testing.T) {
		arr := []string{"you", "shall", "not", "pass"}
		result := TwoSortRefactored(arr)
		expected := "n***o***t"

		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})
}
