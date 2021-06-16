package mystr

import (
	"fmt"
	"testing"
)

// Tests
func TestCat(t *testing.T) {
	x := Cat([]string{"I", "am", "Mila", "the", "cat"})
	if x != "I am Mila the cat" {
		t.Error("Expected: I am Mila the cat, Got:", x)
	}
}

func TestJoin(t *testing.T) {
	x := Join([]string{"I", "am", "Mila", "the", "cat"})
	if x != "I am Mila the cat" {
		t.Error("Expected: I am Mila the cat, Got:", x)
	}
}

// Examples
func ExampleCat() {
	fmt.Println(Cat([]string{"I", "am", "Mila", "the", "cat"}))
	// Output:
	// I am Mila the cat
}

func ExampleJoin() {
	fmt.Println(Join([]string{"I", "am", "Mila", "the", "cat"}))
	// Output:
	// I am Mila the cat
}

// benchmarks
func BenchmarkCat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cat([]string{"I", "am", "Mila", "the", "cat"})
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join([]string{"I", "am", "Mila", "the", "cat"})
	}
}
