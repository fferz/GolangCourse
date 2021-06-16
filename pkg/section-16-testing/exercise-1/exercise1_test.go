package dog

import (
	"fmt"
	"testing"
)

// Tests
func TestYears(t *testing.T) {
	x := Years(10)
	if x != 70 {
		t.Error("Expected: 70, got: ", x)
	}
}

func TestYears2(t *testing.T) {
	x := Years2(10)
	if x != 70 {
		t.Error("Expected: 70, got: ", x)
	}
}

// Examples
func ExampleYears() {
	fmt.Println(Years(5))
	// Output:
	// 35
}

func ExampleYears2() {
	fmt.Println(Years2(5))
	// Output:
	// 35
}

// Benchmarks
func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(5)
	}
}

func BenchmarkYears2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years2(5)
	}
}

/*
- go test
PASS

- go test -bench .
BenchmarkYears-8        1000000000               0.600 ns/op
BenchmarkYears2-8       272709298                4.47 ns/op

- go test -cover
PASS
coverage: 100.0% of statements

- go test -coverprofile coverDog.out

- go tool cover -html=coverDog.out

- godoc -http=:8080
*/
