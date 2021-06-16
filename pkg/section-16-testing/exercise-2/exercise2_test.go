package word

import (
	"fmt"
	"reflect"
	"testing"
)

// tests
func TestUseCounts(t *testing.T) {
	quote := "one two one"

	res := map[string]int{
		"one": 2,
		"two": 1,
	}

	xs := UseCounts(quote)

	eq := reflect.DeepEqual(res, xs)
	if !eq {
		t.Error("Expected: map[one:2, two:1], got: ", xs)
	}

}

func TestCount(t *testing.T) {
	quote := "one two one"
	xs := Count(quote)
	if xs != 3 {
		t.Error("Expected: 3, got: ", xs)
	}
}

// examples
func ExampleUseCounts() {
	quote := "one two one"
	fmt.Println(UseCounts(quote))
	// Output:
	// map[one:2 two:1]
}

func ExampleCount() {
	quote := "one two one"
	fmt.Println(Count(quote))
	// Output:
	// 3
}

// benchmarks
func BenchmarkUseCounts(b *testing.B) {
	quote := "one two one"
	for i := 0; i < b.N; i++ {
		UseCounts(quote)
	}
}

func BenchmarkCount(b *testing.B) {
	quote := "one two one"
	for i := 0; i < b.N; i++ {
		Count(quote)
	}
}

/*
- go test
PASS

- go test -bench .
BenchmarkUseCounts-8     4401966               271 ns/op
BenchmarkCount-8        12867748                92.1 ns/op

- go test -cover
PASS
coverage: 100.0% of statements

- go test -coverprofile coverExercise2.out

- go tool cover -html=coverExercise2.out

- godoc -http=:8080
*/
