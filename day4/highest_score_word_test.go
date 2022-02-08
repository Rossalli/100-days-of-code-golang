package main

import (
	"testing"
)

func TestHighScoreWord(t *testing.T) {

	t.Run("draco dormiens nunquam titillandus", func(t *testing.T) {
		result := HighScoreWord("draco dormiens nunquam titillandus")
		expected := "titillandus"

		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})

	t.Run("aa b", func(t *testing.T) {
		result := HighScoreWord("aa b")
		expected := "aa"

		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})

	t.Run("b aa", func(t *testing.T) {
		result := HighScoreWord("b aa")
		expected := "b"

		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})

	t.Run("i need a taxi up to ubud", func(t *testing.T) {
		result := HighScoreWord("i need a taxi up to ubud")
		expected := "taxi"

		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})
}
