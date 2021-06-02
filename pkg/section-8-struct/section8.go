package section8

import "fmt"

// Struct = Aggregate data type = composite data type = complex data types
// a data structure that allows values of different types
// field types cannot be a pointer type

type person struct {
	first string
	last  string
}
type secretAget struct {
	person // anonymous field - unqualified type name (1)
	ltk    bool
}

func StructData() {
	// we don't say object, we say 'we create a value of type person'
	p1 := person{
		first: "pepa",
		last:  "cow",
	}

	p2 := person{
		first: "andy",
		last:  "jiraffe",
	}
	println("- value of type person")
	fmt.Println("persona 1", p1)
	fmt.Println("nombre p1:", p1.first)
	fmt.Println("persona 2", p2)

	println("- embedded types")
	p3 := secretAget{
		person: person{ // (1) the unqualified type name acts as the field name
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	fmt.Println("secret agent:", p3)
	fmt.Println("secret agent name:", p3.first) // p3.person.first == p3.first (the first field gets promoted)

	println("- anonymous struct")
	// a struct with no name
	p4 := struct {
		first string
		last  string
	}{
		first: "ana",
		last:  "mendez",
	}
	fmt.Println("anonymous struct:", p4)

}
