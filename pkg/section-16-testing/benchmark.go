package section16

import "fmt"

func BanchmarkIntro() {
	Greeting("Ana")
}

func Greeting(s string) string {
	return fmt.Sprint("Welcome ", s)
}
