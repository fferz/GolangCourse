package section9

import "fmt"

func Syntax() {
	// func (r receiver) identifier(parameters) (return(s)) { ... }
	s1 := woo("hola")
	fmt.Println("one return", s1)

	// function with multiple returns
	t1, t2 := buf("uno", "dos")
	fmt.Println("funcion with multiple returns:", t1, t2)
}

func woo(s string) string {
	return fmt.Sprintln("function returns a string:", s)
}

func buf(s string, s1 string) (string, bool) {
	a := fmt.Sprintln("first return", s, s1)
	b := false
	return a, b
}

// ---------------------------------------------------------------------

func VariadicParam() {
	// a function with an unlimited number of parameters
	// the function manage the parameters as a slice
	bubu(1, 2, 3, 4)
}

func bubu(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i := range x {
		sum = sum + x[i]
	}
	fmt.Println("the sum of the params is:", sum)
}

// ---------------------------------------------------------------------

func UnfurlingSlice() {
	/*
		x := []int{1, 2, 3, 4, 5}
		buu(x) // this won't work because buu is expecting int values, not slice
		buu(x...) // this is OK
		buu() // this will work too, because variadic means 0 to unlimited
		func buu(x ...int) {}

		func buu(s string, x ...int) {} // OK
		func buu(x ...int, s string) {} // won't work

	*/
}

// ---------------------------------------------------------------------

func Defer() {
	// will defer the execution of a function until the surrounding function returns

	defer foo1()
	foo2()
	// prints:
	// foo2
	// foo1
}

func foo1() {
	fmt.Println("foo1")
}

func foo2() {
	fmt.Println("foo2")
}

// ---------------------------------------------------------------------

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func Methods() {

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak() // any value of type 'secretAgent' has access to the 'speak' method

	// func (r receiver) identifier(parameters) (return(s)) { ... }
	p1 := person{
		first: "lisa",
		last:  "simpson",
	}
	fmt.Println("person does not have access to the speak method. \n", p1)
	// p1 does not have access to speak method.
}

func (s secretAgent) speak() {
	//the receiver in a function is going to attach this method to any value of the receiver type
	// any value of type secret agent has access to this method
	fmt.Println("I am a secret agent")
}

// ---------------------------------------------------------------------

type human interface {
	speak()
}

func InterfacesAndPolymorphism() {
	sa2 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	fmt.Printf("%T\n", sa2)
	// sa is of type secretAgent but also, of type human, because the method speak is of type human
	humanfunc(sa2) // I can call this function because sa2 is also type human

	p2 := person{
		first: "lisa",
		last:  "simpson",
	}
	fmt.Println("humanfunc won't accept a type person ", p2)
	//humanfunc(p2) // error, person has not have the speak method

	println("assertion")
	// x.(T)
	// asserts that x is not nil and that the value stored in x is of type T

	println("empty interface")
	/*
		type human interface {} // this is an empty interface
		all the types implements an empty interface
		now this will work => humanfunc(p2)

		A type implements any interface comprising any subset of its methods
		and may therefore implement several distinct interfaces. For instance,
		all types implement the empty interface
	*/
}

func humanfunc(h human) {
	println("humanfunc")
	//h.speak()
	fmt.Println("I am a human.", h)
	// h.ltk does not work, but if I use assertion, it will work: h.(secretAgent).first
	fmt.Println(h.(secretAgent).person.first)
	// I assert that h is not nil and that it is of type secretAgent
}

// ---------------------------------------------------------------------

func AnonymousFunc() {
	// these functions has no identifier
	func() {
		fmt.Println("I'm an anonymous function.")
	}()
	func(x int) {
		fmt.Println("I'm an anonymous function with params:", x)
	}(22)
}

// ---------------------------------------------------------------------

func FuncExpression() {
	// we assign a function to a variable
	fmt.Println("function expression")
	f := func() {
		fmt.Println("I am a function assigned to a variable")
	}
	f()
	g := func(x int) {
		fmt.Println("I am a function assigned to a variable with params: ", x)
	}
	g(22)
}

// ---------------------------------------------------------------------

func ReturningFun() {
	fmt.Println("returning function")
	x := myFunction()
	fmt.Printf("%T\n", x)       // type of x => func() int
	fmt.Println(x)              // returns 0x109ecb0
	fmt.Println(x())            // returns 22
	fmt.Println(myFunction()()) // returns 22
}

func myFunction() func() int {
	// return type = func() int
	return func() int {
		return 22
	}
}

// ---------------------------------------------------------------------

func Callbacks() {
	// when I pass a function as an argument to another function
	fmt.Println("callbacks")
	p := []int{1, 2, 3, 4, 5, 6}
	result := sum(p...)
	fmt.Println("function sum result: ", result)
	r := evenSum(sum, p...)
	fmt.Println("result of even sum: ", r)
}

func sum(x ...int) int {
	fmt.Println("sum function")
	fmt.Printf("params type: %T\n", x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func evenSum(f func(x ...int) int, y ...int) int {
	// parameters:
	// 1. identifier = f; type = funt(x ...int) int
	// 2. identifier = y; type = ...int
	var args1 []int
	for _, v := range y {
		if v%int(2) == 0 {
			args1 = append(args1, v)
		}
	}
	args2 := f(args1...)
	return args2
}

// ---------------------------------------------------------------------

func Closures() {
	// limited scope of variables

}

// ---------------------------------------------------------------------

func Recursion() {
	/*
		func factorial(n int) int {
			if n == 0 { return 1 }
			return n * factorial (n -1)
		}

		func loop(n int) int {
			res := 1
			for i, v := range n {
				if i > 0 {
					res *= v
				}
			}
			for _; n > 0; n-- {
				res *= n
			}
			return res
		}

	*/
}

// ---------------------------------------------------------------------

// eerything in Go is passed by value

// a value can be of more than one type

// funtions are first class citizens in Go
