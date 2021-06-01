package section6

import "fmt"

func Loop() {
	print("Loops")
	// there is no while in Go
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}

func NestingLoops() {
	println("nesting loops")
	for i := 0; i <= 5; i++ {
		fmt.Printf("outer loop %v \n", i)
		for j := 0; j <= 2; j++ {
			fmt.Printf("\t the inner loop %v \n", i)
		}
	}
}

func ForStatement() {
	println("for statement")
	/*
		1. for init; condition; post {}
		2. for condition {}
		3. for {}
	*/
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	for {
		if x < 5 {
			fmt.Println(x)
			x++
		} else {
			break
		}
	}
}

func BreakAndContinue() {
	/*

		x := 1
		for {
			if x < 10 {
				break
			}
			if x % 2 != 0 {
				// do nothing, skip this step
				continue
			}
			fmt.Println(x)
		}

	*/
}

func Exercise1() {
	println("numbers from 0 to 200, and print them also turn into letters")
	i := 0

	for i <= 200 {
		//fmt.Printf("number: %d, letter: %c \n", i, i)
		// number: 192, letter: À
		//fmt.Printf("number: %d, letter: %#U \n", i, i)
		// number: 65, letter: U+0041 'A'
		fmt.Printf("number: %d, hexa: %#X letter: %#U \n", i, i, i)
		i++
	}
}

func IfStatement() {
	/*
		if condition {}
		if x := 42; x == 42 {
			// we can use semicolon to have 2 statements in one line
		}
		if condition {

		} else if condition {

		} else {}
	*/
	if true {
		println("hello")
	}
}

func Exercise2() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func SwitchStatement() {

	switch {
	case false:
		fmt.Println("case 1")
	case true:
		fmt.Println("case 2")
		fallthrough //  print and continue evalutaing cases
	case true:
		fmt.Println("case 3")
	case false:
		fmt.Println("case 4")
		fallthrough // nothing happens here
	default:
		fmt.Println("default case")
	}

}

/*
	Conditional logic operators

	&& = and
	|| = or
	! = not


*/
