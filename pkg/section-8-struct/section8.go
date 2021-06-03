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

type person1 struct {
	first_name          string
	last_name           string
	favourite_ice_cream []string
}

func Exercise1() {
	p1 := person1{
		first_name:          "ana",
		last_name:           "perez",
		favourite_ice_cream: []string{"vainilla", "chocolate", "tramontana"},
	}

	p2 := person1{
		first_name:          "laila",
		last_name:           "pez",
		favourite_ice_cream: []string{"anana", "lemon", "orange"},
	}

	fmt.Println("first name: ", p1.first_name)
	fmt.Println("last_name: ", p1.last_name)
	for k, v := range p1.favourite_ice_cream {
		fmt.Printf("key: %v, value: %v \n", k, v)

	}
	fmt.Println(p2)

	m1 := map[string]person1{
		p1.last_name: p1,
		p2.last_name: p2,
	}
	fmt.Println("map ", m1[p1.last_name].first_name)

	for _, v1 := range m1 {
		fmt.Printf(v1.first_name)
		fmt.Printf(v1.last_name)
		for k1, v1 := range v1.favourite_ice_cream {
			fmt.Printf("key: %v, value: %v \n", k1, v1)
		}
	}

}

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func Exercise2() {
	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}
	s1 := sedan{
		vehicle: vehicle{
			doors: 5,
			color: "black",
		},
		luxury: false,
	}
	fmt.Println("truck: ", t1)
	fmt.Println("sedan: ", s1)
}

func Exercise3() {
	an := struct {
		animal string
		color  string
	}{
		animal: "cat",
		color:  "white",
	}
	fmt.Println(an)
	an1 := struct {
		colors map[int]string
		car    vehicle
		things []string
	}{
		colors: map[int]string{
			1: "pink",
			2: "blue",
		},
		car: vehicle{
			doors: 4,
			color: "white",
		},
		things: []string{"one", "two", "three"},
	}
	fmt.Println("an2", an1)
	println("loop for printing an2")
	for i, v := range an1.colors {
		fmt.Printf("key: %v, value: %v\n", i, v)
	}
	fmt.Println("an2 car: ", an1.car)
	for i1, v1 := range an1.things {
		fmt.Printf("things key: %v, value: %v\n", i1, v1)
	}
}
