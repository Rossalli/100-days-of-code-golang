package main

import "testing"

func TestCompArrays(t *testing.T) {

	t.Run("a = [121, 144, 19, 161, 19, 144, 19, 11]  and b = [121, 14641, 20736, 361, 25921, 361, 20736, 361] it's same", func(t *testing.T) {
		same := Comp([]int{121, 144, 19, 161, 19, 144, 19, 11}, []int{121, 14641, 20736, 361, 25921, 361, 20736, 361})

		if !same {
			t.Errorf("Expected result is true")
		}
	})

	t.Run("a = [121, 144, 19, 161, 19, 144, 19, 11]  and b = [132, 14641, 20736, 361, 25921, 361, 20736, 361] it's not the same", func(t *testing.T) {
		same := Comp([]int{121, 144, 19, 161, 19, 144, 19, 11}, []int{132, 14641, 20736, 361, 25921, 361, 20736, 361})
		if same {
			t.Errorf("Expected result is false")
		}
	})

	t.Run("a = [121, 144, 19, 161, 19, 144, 19, 11]  and b = [11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19] it's not the same", func(t *testing.T) {
		same := Comp([]int{121, 144, 19, 161, 19, 144, 19, 11}, []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19})
		if !same {
			t.Errorf("Expected result is false")
		}
	})

	t.Run("a = [121, -144, 19, -161, 19, -144, 19, -11]  and b = [11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19] it's not the same", func(t *testing.T) {
		same := Comp([]int{-121, -144, 19, -161, 19, -144, 19, -11}, []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19})
		if !same {
			t.Errorf("Expected result is true")
		}
	})

	t.Run("a = [2, 3, 3]  and b = nil it's not the same", func(t *testing.T) {
		same := Comp([]int{2, 3}, nil)
		if same {
			t.Errorf("Expected result is true")
		}
	})

}
