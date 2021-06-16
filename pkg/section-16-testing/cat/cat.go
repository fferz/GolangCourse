package section16

import (
	"fmt"
	"strings"

	"github.com/fferz/GolangCourse/pkg/section-16-testing/cat/mystr"
)

const s = "I am Mila the cat"

func MilaCat() {
	xs := strings.Split(s, "")

	for _, v := range xs {
		fmt.Println(v)
	}

	fmt.Println(mystr.Cat(xs))
	fmt.Println(mystr.Join(xs))
}
