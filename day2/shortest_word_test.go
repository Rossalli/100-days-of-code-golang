package main

import (
	"testing"
)

func TestShortestWordTest(t *testing.T) {

	phrase1 := "The things you used to own, now they own you"
	t.Run("Shortest word in phrase has lengeth 2", func(t *testing.T) {
		result := ShortestWord(phrase1)
		expected := 2
		if result != expected {
			t.Errorf("result '%d', expected '%d'", result, expected)
		}
	})

	phrase2 := "The lower you fall, the higher you'll fly"
	t.Run("Shortest word in phrase has lengeth 3", func(t *testing.T) {
		result := ShortestWord(phrase2)
		expected := 3
		if result != expected {
			t.Errorf("result '%d', expected '%d'", result, expected)
		}
	})
}
