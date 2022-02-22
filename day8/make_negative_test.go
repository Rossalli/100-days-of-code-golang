package main

import (
	"testing"
)

func TestMakeNegative(t *testing.T) {

	t.Run("n = 0 return 0", func(t *testing.T) {
		result := MakeNegative(0)
		expected := 0
		if result != expected {
			t.Errorf("result '%d', expected '%d'", result, expected)
		}
	})

	t.Run("n = 7 return -7", func(t *testing.T) {
		result := MakeNegative(7)
		expected := -7
		if result != expected {
			t.Errorf("result '%d, expected '%d'", result, expected)
		}
	})

	t.Run("n = -42 return -42", func(t *testing.T) {
		result := MakeNegative(-42)
		expected := -42
		if result != expected {
			t.Errorf("result '%d', expected '%d'", result, expected)
		}
	})

}
