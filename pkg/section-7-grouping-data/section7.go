package section7

import "fmt"

func Arrays() {
	var x [5]int
	fmt.Println(x)
	x[3] = 12
	fmt.Println(x)
	// len(x)
	// the len is part of an array's type => type [5]int != type [4]int
	// arrays are not really used in Go
	// use slices!
}

func Slices() {
	// composite literal =>   x := type{values}
	// this is the way to construct values for strucs, arrays, slices and maps
	// they create a new value each time they're evaluated
	// slices allows us to group together values of the same type
	a := []int{1, 2, 3, 4}
	fmt.Println(a)
	fmt.Println("length", len(a))
	fmt.Println("index position 2", a[2]) // access by index
	// loop an array, slice, string or map => we use 'range'
	println("- loop using range")
	for i, v := range a {
		fmt.Printf("index: %v, value: %v \n", i, v)
	}

	// slicing a slice
	println("- slicing a slice")
	fmt.Println("from position 1:", a[1:])
	fmt.Println("until position 2:", a[:2])
	fmt.Println("from position 1 to 3 (not included)", a[1:3])

	println("- for loop")
	for i := 0; i < len(a); i++ {
		fmt.Printf("index: %v, value: %v \n", i, a[i])
	}

	println("- append values to slice")
	b := append(a, 10, 11, 12) // append extra values to a slice
	fmt.Println(b)
	c := []int{100, 200, 300, 400, 500}
	d := append(a, c...) // append two slices
	fmt.Println(d)

	println("- delete from a slice")
	e := append(d[2:4], d[6:8]...)
	fmt.Println(e)

	println("- create a slice using make")
	f := make([]int, 5, 10) // allocates an underlying array of size 10, and returns a slice of lenght 5 (with a capacity of 10)
	fmt.Println(f)
	fmt.Printf("size: %v, capacity: %v \n", len(f), cap(f))
	println("append 9 items to array")
	g := append(f, d...)
	fmt.Println(g)
	fmt.Printf("size: %v, capacity: %v \n", len(g), cap(g))

	println("- multidimensional slice")
	h := []string{"one", "two", "three"}
	i := []string{"a", "b", "c"}
	hi := [][]string{h, i}
	fmt.Println("multidimensional array:", hi)
}

func Maps() {
	println("Maps")
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(m)
	fmt.Println(m["one"])  // => 1
	fmt.Println(m["five"]) // => 0 because this does not exists

	println(" - check if a key is valid")
	v, ok := m["five"]
	if ok == true {
		fmt.Println("key is valid, value:", v)
	}
	// another way to write it
	if v, ok := m["five"]; ok {
		fmt.Println("key is valid, value:", v)
	} else {
		fmt.Println("key is not valid, value:", v)
	}

	println("- add elements to map")
	m["four"] = 4
	m["five"] = 5
	for k, v := range m {
		fmt.Printf("key: %v, value: %v \n", k, v)
	}

	println("- delete")
	delete(m, "five")
	println(" delete five key")
	for k, v := range m {
		fmt.Printf("key: %v, value: %v \n", k, v)
	}
	println("- delete a key that does not exists") // there are not any errors
	delete(m, "six")
	// so we need to check if the key exists before we delete it!
	if v, ok := m["six"]; ok {
		delete(m, "six")
	} else {
		fmt.Println("key does not exists", v)
	}
}
