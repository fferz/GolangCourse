package pointers

import "fmt"

func HelloPointers() {
	a := 12
	fmt.Println("a:", a)
	fmt.Println("a address:", &a)
	fmt.Println("types")
	fmt.Printf("a = %T\n", a)
	fmt.Printf("&a = %T\n", &a) // this type is a pointer

	// operator & returns address of the variable
	// &a => address => type pointer *int (*int != int)
	// var b *int := &a

	// operator * returns the value stored in the address
	// b:= &a => b has the address of a
	// *b     => now I can get the value stored in the address (derefence of b)

	// *&a => give me the address of a, and then, derefence it so I could get the value of a

	fmt.Println("a value:", a)
	var b *int = &a
	fmt.Println("var b *int := &a, so b = ", b)
	*b = 13
	fmt.Println("*b = 13")
	fmt.Println("now a = 13:", a)
}

// everything in go is passed by value

