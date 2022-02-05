package main

import (
	"testing"
)

func TestMakeUpperCase(t *testing.T) {

	t.Run("Make uppercase string abcd", func(t *testing.T) {
		result := MakeUpperCase("abcd")
		expected := "ABCD"
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}

	})

	t.Run("Make uppercase string ç1qAZ", func(t *testing.T) {
		result := MakeUpperCase("ç1qAZ")
		expected := "Ç1QAZ"
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}	

	})
}
