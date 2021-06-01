/*

rune = int32
byte = uint8
*/

package section5

import (
	"fmt"
	"runtime"
)

var a bool

func VariablesTypes() {

	println("Bool type")
	fmt.Println(a)
	fmt.Println(7 == 17)
}

func VariablesEnv() {
	println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

const (
	b     = 42
	c     = "James Bond"
	d     = 42.78
	e int = 10
)

func VariablesConst() {
	println("Constants")
	println(b)
	fmt.Printf("%T\n", b)
	println(b)
	println(c)
	fmt.Printf("%T\n", c)
	println(d)
	fmt.Printf("%T\n", d)
	// typed const e cannot be assign to untyped const e
	// e = d <= this throws an error
	fmt.Println(e)
}

func VariableIOTA() {
	println("Iota")
	// pre declared identified
	// successive untyped integer constants
	const (
		a = iota
		b
		c
	)
	println("a = 0,", a)
	println("b = 1,", b)
	println("c = 2,", c)

	// iota use
	/*
		const (
			_ = iota =>  0
			kb = 1 << (iota * 10) => 10000000000 (iota = 1)
			mb = 1 << (iota * 10) => 100000000000000000000 (iota = 2)
			gb = 1 << (iota * 10) => 1000000000000000000000000000000 (iota = 3)
		)
	*/
}

func Exercise1() {
	println("prints a number in dec, bin and hex")
	d := 23
	fmt.Printf("%d\t%b\t%#x", d, d, d)
}

func Exercise2() {
	println("comparissons")
	e := (13 == 13)
	f := (13 <= 13)
	g := (13 >= 13)
	h := (13 != 13)
	i := (13 < 13)
	j := (13 > 13)
	fmt.Println("e", e)
	fmt.Println("f", f)
	fmt.Println("g", g)
	fmt.Println("h", h)
	fmt.Println("i", i)
	fmt.Println("j", j)

}

func Exercise3() {
	println("typed and untyped constants")
	const (
		h        = "hola"
		i        = 1
		j string = "hi"
		k int    = 2
	)
}

func Exercise4() {
	l := 12
	fmt.Printf("%d\t%b\t%x\n", l, l, l)
	m := l << 1
	fmt.Printf("%d\t%b\t%x\n", m, m, m)
}

func Exercise5() {
	println("raw string literal")
	n := `hola
	como estas
	?`
	fmt.Println(n)
}

func Exercise6() {
	const (
		o = 2019 + iota
		p = 2017 + iota
		q = 2017 + iota
		r = 2017 + iota
	)
	fmt.Println(o, p, q, r)
}
