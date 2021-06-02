/*

- function(nums ...int){}
this means 'undefined number of int numbers as parameters'

- type 'interface{}'
every value is also of type interface{}
the parameter '...interface{}' means as many arguments of any type as you wish. this is called 'Variadic Parameter'

- underscore character _
we use it to throw away returns


- short declaration operator ':='
it declares a variable and also assign a value to it, all at the same time
this works inside a function body
a := 2
later, if we want to assign a new value, we just use '='
a = 3


- format code
go fmt

- variable declarations
var y = 13 // outside function scope
e := 22    // inside function scope (use this)

var z int // outside function
here I'm declaring a variable with the identifier 'z'.
this variable is not initialized, so the compiler assigns the ZERO VALUE of type int to 'z'.

- zero values (or default values)
boolean : false
int : 0
float: 0.0
string: ""
pointers, interfaces, functions, slices, channers, maps: nil

- static programming language

a := "hola"
a := `hello`

and this will throw an error, because a was declared as a string variable,
we cannot change it to int
	a := "hello"
	fmt.Println(`a := "hello"`)
	a = 1

- create my own type

var a int
a = 42
type hotdog int
var b hotdog
b = 50

a = b won't work, because those are different types
// this is called type alias and it is not a good practise

- conversion not casting

var a int
a = 42
type hotdog int
var b hotdog
b = 50
a = int(b) <= this is CONVERSION

is casting, but in Go we call it Conversion.

*/

package section4

import "fmt"

func Exercise1() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Print(`
		y,
		z
	`)
}

var x int
var y string
var z bool

func Exercise2() {
	println("zero values")
	println("exercise2, x:", x)
	println("exercise2, y:", y)
	println("exercise2, z:", z)
}

var a int = 42
var b string = "James Bond"
var c bool = true

func Exercise3() {
	a1 := fmt.Sprintf("%d", a)
	b1 := fmt.Sprintf("%s", b)
	c1 := fmt.Sprintf("%t", c)
	d := fmt.Sprintf("%v\t%v\t%v", a, b, c)
	println("exercise3, a:", a1)
	println("exercise3, b:", b1)
	println("exercise3, c:", c1)
	println("exercise3, d:", d)
}

type myType int

var e myType

func Exercise4() {
	println("my own type")

	println("e value:", e)
	fmt.Printf("%T\n", e)
	e = 42
	println("e = 42, e:", e)
}

// f type is the underlying type of hotdog, which is int
var f int

func Exercise5() {
	f := int(e)
	println("exercise 5, value of f:", f)
}

func Exercise6() {

}
