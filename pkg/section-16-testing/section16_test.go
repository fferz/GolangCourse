package section16

import (
	"testing"
)

func TestMySum(t *testing.T) {
	x := mySum(2, 3)
	if x != 5 {
		t.Error("Expected 5, got: ", x)
	}

}

// Table Tests
func TestMySum2(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{
			data:   []int{2, 2, 2},
			answer: 6,
		},
		{
			data:   []int{-1, 2, 6},
			answer: 7,
		},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Errorf("Expected: %v, got: %v", v.answer, x)
		}
	}

}
