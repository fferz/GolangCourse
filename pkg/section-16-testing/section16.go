// Package section16 is about Testing.
package section16

import (
	"fmt"

	word "github.com/fferz/GolangCourse/pkg/section-16-testing/exercise-2"
	quote "github.com/fferz/GolangCourse/pkg/section-16-testing/exercise-2-quote"
)

/*
Testing

- we put tests in a file that ends with: '_test.go'
- the tests file must be in the same package as te one being tested
- a test is a func with this name: 'func TestMyFunction(t *testing.T)'
- we run tests using: go test
*/

/*
package math

import "testing"

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1,2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}
*/

func Example1() {
	m := mySum(1, 2, 3)
	fmt.Println("sum(1, 2, 3) = ", m)
}

// mySum adds an unlimited number of values of type int
func mySum(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}

// an example is a great way to document the code, and to test it, all at the same time
// test as documentation example (go test) - testing package
func ExampleMySum() {
	fmt.Println(mySum(2, 3))
	// Output:
	// 5
}

/*
Golint

- gofmt => formats go code
- go vet => repots suspicious constructs
- golint => reports poor coding style

golint install: go get -u golang.org/x/lint/golint
golint usage: go lint ./...
*/

func Exercise2() {
	fmt.Println("amout of use of each word")
	fmt.Println(word.UseCounts(quote.SerenaWilliams))
	fmt.Println("total words")
	fmt.Println(word.Count(quote.SerenaWilliams))

}
