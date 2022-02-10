package main

import "testing"

func TestAccum(t *testing.T) {

	t.Run("abcd -> A-Bb-Ccc-Dddd", func(t *testing.T) {
		result := Accum("abcd")
		expected := "A-Bb-Ccc-Dddd"
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})

	t.Run("Wololo -> W-Oo-Lll-Oooo-Lllll-Oooooo", func(t *testing.T) {
		result := Accum("Wololo")
		expected := "W-Oo-Lll-Oooo-Lllll-Oooooo"
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})

	t.Run("kamehameha -> K-Aa-Mmm-Eeee-Hhhhh-Aaaaaa-Mmmmmmm-Eeeeeeee-Hhhhhhhhh-Aaaaaaaaaa", func(t *testing.T) {
		result := Accum("kamehameha")
		expected := "K-Aa-Mmm-Eeee-Hhhhh-Aaaaaa-Mmmmmmm-Eeeeeeee-Hhhhhhhhh-Aaaaaaaaaa"
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})

	t.Run("FusRoDah -> F-Uu-Sss-Rrrr-Ooooo-Dddddd-Aaaaaaa-Hhhhhhhh", func(t *testing.T) {
		result := Accum("FusRoDah")
		expected := "F-Uu-Sss-Rrrr-Ooooo-Dddddd-Aaaaaaa-Hhhhhhhh"
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})

}

func TestAccumRefactored(t *testing.T) {

	t.Run("abcd -> A-Bb-Ccc-Dddd", func(t *testing.T) {
		result := AccumRefactored("abcd")
		expected := "A-Bb-Ccc-Dddd"
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})

	t.Run("Wololo -> W-Oo-Lll-Oooo-Lllll-Oooooo", func(t *testing.T) {
		result := AccumRefactored("Wololo")
		expected := "W-Oo-Lll-Oooo-Lllll-Oooooo"
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})

	t.Run("kamehameha -> K-Aa-Mmm-Eeee-Hhhhh-Aaaaaa-Mmmmmmm-Eeeeeeee-Hhhhhhhhh-Aaaaaaaaaa", func(t *testing.T) {
		result := AccumRefactored("kamehameha")
		expected := "K-Aa-Mmm-Eeee-Hhhhh-Aaaaaa-Mmmmmmm-Eeeeeeee-Hhhhhhhhh-Aaaaaaaaaa"
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})

	t.Run("FusRoDah -> F-Uu-Sss-Rrrr-Ooooo-Dddddd-Aaaaaaa-Hhhhhhhh", func(t *testing.T) {
		result := AccumRefactored("FusRoDah")
		expected := "F-Uu-Sss-Rrrr-Ooooo-Dddddd-Aaaaaaa-Hhhhhhhh"
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})

}
