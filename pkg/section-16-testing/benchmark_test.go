package section16

import (
	"fmt"
	"testing"
)

func TestGreeting(t *testing.T) {
	s := Greeting("Anna")
	if s != "Welcome Anna" {
		t.Error("expected: Welcome Anna, got: ", s)
	}
}

func ExampleGreeting() {
	fmt.Println(Greeting("Anna"))
	// Output:
	// Welcome Anna
}

// benchmark will run this code a lot of times until it gets a
// statiscal accurate measure of the time this code needs to run
func BenchmarkGreeting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greeting("Anna")
	}
}

/*
go test -bench .
go test -bench Greeting

BenchmarkGreeting-8      8731932               131 ns/op
the test run on 8 cores, it runed 8731932 times, and it took 131 ns per operation.
*/

/*
Coverage

go test -cover
PASS
coverage: 55.6% of statements

go test -coverprofile coverprofile.out
(this generates a file with the cover data)

go tool cover -html=coverprofile.out
(this shows a page in the browser highlighted with the cover and uncover parts)

*/
