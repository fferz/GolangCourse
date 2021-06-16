package section14

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ErrorIntro() {
	n, err := fmt.Println("hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}

// scans input from keyboard
func ScanExample() {
	var answer1, answer2, answer3 string

	fmt.Print("name: ")
	_, err := fmt.Scan((&answer1))
	if err != nil {
		panic(err)
	}

	fmt.Print("color: ")
	_, err2 := fmt.Scan(&answer2)
	if err2 != nil {
		panic(err2)
	}

	fmt.Print("song: ")
	_, err3 := fmt.Scan(&answer3)
	if err3 != nil {
		panic(err3)
	}

	fmt.Println(answer1, answer2, answer3)

}

func ReaderExample() {
	f, err := os.Create("OneFile.txt") // file's name
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("hello") // something to write in file

	io.Copy(f, r)
}

func ReaderExample2() {
	f, err := os.Open("OneFile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(r))
}

/*
how to print out the errors?

- fmt.Println()
- log.Println()
- log.Fatalln()
	. os.Exit()
- log.Panicln()
	. deferred functions run
	. can use 'recover'
- panic()

which to use? with log we have more options and we could write it out in a new file.
We should use
	log.println for normal logging,
	log.falatln : the entire program is going to shut down immediately and no defer functions are run
	log.panic : it will run defer functions, and walk trough the call stack in reverse order and we could then use recover
*/

func ErrorShow() {
	println("error show")
	defer foo()
	_, err := os.Open("OneFile1.txt")
	if err != nil {
		log.Panicln(err)
		// fmt.Println(err) => open OneFile1.txt: no such file or directory | [executes foo at last]

		// log.Println(err) => 2021/06/15 11:37:03 open OneFile1.txt: no such file or directory | [executes foo at last]

		// log.Fatalln(err) => 2021/06/15 11:37:50 open OneFile1.txt: no such file or directory | [does not execute foo]
		//										exit status 1 (code zero indicates success, non-zero an error - package os, func exit())

		// log.Panicln(err) => | [executes foo at first]
		/*
					2021/06/15 11:52:10 open OneFile1.txt: no such file or directory
			function foo executed
			panic: open OneFile1.txt: no such file or directory


			goroutine 1 [running]:
			log.Panicln(0xc000098f50, 0x1, 0x1)
			        /usr/local/go/src/log/log.go:365 +0xac
			github.com/fferz/GolangCourse/pkg/section-14-error-handling.ErrorShow()
			        /Users/fernandafernandez/proyectos/golang course/pkg/section-14-error-handling/section14.go:94 +0xb5
			main.main()
			        /Users/fernandafernandez/proyectos/golang course/main.go:106 +0x20
			exit status 2
		*/

		// panic(err)      => | [executes foo at first]
		/*
					panic: open OneFile1.txt: no such file or directory

			goroutine 1 [running]:
			github.com/fferz/GolangCourse/pkg/section-14-error-handling.ErrorShow()
			        /Users/fernandafernandez/proyectos/golang course/pkg/section-14-error-handling/section14.go:96 +0x74
			main.main()
			        /Users/fernandafernandez/proyectos/golang course/main.go:106 +0x20
			exit status 2

		*/
	}
}

func foo() {
	fmt.Println("function foo executed")
}

func CreateLogFile() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		log.Println(err)
	}
	defer f2.Close()

	fmt.Println("check the log.txt file in the directory")
}

// Recover after panic
func MainRecover() {
	f()                                       // 1
	fmt.Println("Returned normally from f()") //19
}

/*
2. Calling g()
4. Printing in g 0
6. Printing in g 1
8. Printing in g 2
10. Printing in g 3
13. Panick!!
15. Defer in g 3
15. Defer in g 2
15. Defer in g 1
15. Defer in g 0
18. Recovered in f 4
19. Returned normally from f()
*/

func f() {
	defer func() { //16
		if r := recover(); r != nil { //17
			fmt.Println("Recovered in f", r) //18
		}
	}()
	fmt.Println("Calling g()")                //2
	g(0)                                      //3
	fmt.Println("Returned normally from g()") // this never prints
}

func g(i int) {
	if i > 3 { //12
		fmt.Println("Panick!!")     //13
		panic(fmt.Sprintf("%v", i)) //14
	}
	defer fmt.Println("Defer in g", i) //15
	fmt.Println("Printing in g", i)    // 4 - 6 - 8 - 10
	g(i + 1)                           //5 - 7 - 9 - 11
}

func ErrorsWithInfo() {
	// package errors - func New()
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}

}

var SquareError = errors.New("Square root of negative number.")

func sqrt(f float64) (float64, error) {
	SquareError2 := fmt.Errorf("Square root of negative number: %v", f)
	fmt.Printf("error type %T \n ", SquareError) //*errors.errorString
	if f < 0 {
		// return 0, SquareError
		// return 0, errors.New("Square root of negative number.")
		// return 0, fmt.Errorf("Square root of negative number: %v", f)
		// return 0, SquareError2
		return 0, mathError{f, SquareError2}
	}
	return f, nil
}

/*
Create error messages:

errors.New() -> error type
fmt.Errorf() -> error type (allows to pass a variable to the error)
fmt.Sprintf() -> string type

*/

type mathError struct {
	num float64
	err error
}

func (m mathError) Error() string {
	return fmt.Sprintf("A math error ocurred. Value: %v, Error: %v ", m.num, m.err)
}

// exercise 1
type person1 struct {
	First   string
	Last    string
	Sayings []string
}

func Exercise1() {
	p1 := person1{
		First:   "Ana",
		Last:    "Kournikova",
		Sayings: []string{"Hello", "Hi"},
	}

	bs, err := json.Marshal(&p1)
	if err != nil {
		log.Fatalln("JSON did not marshal", err)
	}
	fmt.Println(string(bs))
}

// Exercise 2
func Exercise2() {
	println("use fmt.Errorf()")
	res, err := sum(5, -2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}

func sum(a, b int) (int, error) {
	if a < 0 {
		return 0, fmt.Errorf("value must be positive: %v ", a)
	}
	if b < 0 {
		return 0, fmt.Errorf("value must be positive: %v", b)
	}
	return a + b, nil
}

// exercise 3
type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error %v", ce.info)
}

func foo1(e error) {
	fmt.Println("foo run, error:", e)
}
func Exercise3() {
	ce := customErr{
		info: "This is my error",
	}
	foo1(ce)
}

// exercise 4
func Exercise4() {
	// return 0, sqrtError{"50.2289N", "99.4658W", "hubo un error"}
}
