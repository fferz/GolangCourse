package section11

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func JsonMarshal() {
	println("Marshal")

	p1 := person{
		First: "lara",
		Last:  "perez",
		Age:   29,
	}
	p2 := person{
		First: "mara",
		Last:  "mez",
		Age:   22,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	b, err := json.Marshal(people)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

}

func JsonUnmarshal() {
	println("Unmarshal")

	peopleJson := []byte(`[{"First":"lara",  "Last":"perez",  "Age":29},{"First":"mara",  "Last":"mez",   "Age":22}]`)

	var people []person
	err := json.Unmarshal(peopleJson, &people)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v \n", people)
}

// sort
func Sorting() {
	x1 := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	x2 := []string{"one", "two", "three", "four"}

	println("sort int")
	fmt.Println(x1)
	x1r := sort.IntsAreSorted(x1)
	if x1r == true {
		fmt.Println("slice sorted")
	} else {
		sort.Ints(x1)
	}
	fmt.Println(x1)

	println("sort string")
	fmt.Println(x2)
	x2r := sort.StringsAreSorted(x2)
	if x2r == true {
		fmt.Println("slice sorted")
	} else {
		sort.Strings(x2)
	}
	fmt.Println(x2)

	fmt.Println("orden alfabetico: ", "Zaila" < "Pedro")
}

// sort advanced
type ByName []person

// implementation of Interface interface
func (n ByName) Len() int           { return len(n) }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool { return n[i].First < n[j].First }

func SortAdvanced() {
	people := []person{
		{First: "Mauro", Last: "Morla", Age: 35},
		{First: "Zara", Last: "Benitez", Age: 35},
		{First: "Ana", Last: "Perez", Age: 32},
	}

	fmt.Println(people)

	// sort by first name
	sort.Sort(ByName(people))
	fmt.Println("sorted by name: ", people)
}

// bcrypt
func Encripting() {
	pass := []byte(`password`)
	r, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("password's hash: ", string(r))
	}

	pass1 := []byte(`pass123`)
	err1 := bcrypt.CompareHashAndPassword(r, pass1)
	fmt.Println("compare pass = 'pass123' with hash: ")
	if err1 != nil {
		fmt.Println("wrong password")
	} else {
		fmt.Println("success!")
	}

}
