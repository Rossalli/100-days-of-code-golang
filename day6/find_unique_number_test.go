package main

import "testing"

func TestFindUniqueNumber(t *testing.T) {

	t.Run("1, 1, 42, 0, 0", func(t *testing.T) {
		result := FindUniqueNumber([]float32{1, 1, 42, 0, 0})
		expected := float32(42)
		if result != expected {
			t.Errorf("result '%f', expected '%f'", result, expected)
		}
	})

	t.Run("1, 1, 0.5, 1, 1", func(t *testing.T) {
		result := FindUniqueNumber([]float32{0, 0, 0.5, 0, 0})
		expected := float32(0.5)
		if result != expected {
			t.Errorf("result '%f', expected '%f'", result, expected)
		}
	})
}
