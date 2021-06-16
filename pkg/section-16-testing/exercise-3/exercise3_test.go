package myaverage

import (
	"fmt"
	"testing"
)

// Tests
func TestCenteredAvg(t *testing.T) {
	x := []int{1, 2, 3, 4}
	y := CenteredAvg(x)
	z := 2.5
	if y != z {
		t.Error("Expected: 2,5, got: y")
	}
}

func TestCenteredAvg2(t *testing.T) {

	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		{
			data:   []int{2, 2, 2, 2},
			answer: 2,
		},
		{
			data:   []int{2, 5, 15, 10, 20},
			answer: 10,
		},
	}

	for _, v := range tests {
		x := CenteredAvg(v.data)
		if x != v.answer {
			t.Errorf("Expected: %v, got: %v", v.answer, x)
		}
	}
}

// Examples
func ExampleCenteredAvg() {
	x := []int{1, 2, 3, 4}
	fmt.Println(CenteredAvg(x))
	// Output:
	// 2.5
}

// Benchmarks
func BenchmarkCenteredAvg(b *testing.B) {
	x := []int{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		CenteredAvg(x)
	}

}

/*
- go test
PASS

- go test -bench .
BenchmarkCenteredAvg-8          16881387                69.1 ns/op

- go test -cover
PASS
coverage: 100.0% of statements

- go test -coverprofile coverExercise3.out

- go tool cover -html=coverExercise3.out

- godoc -http=:8080
*/
