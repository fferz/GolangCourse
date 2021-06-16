package section12

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

// definition => func (wg *WaitGroup) Add(delta int)
/*
the method set of a value type is a
subset of the method set of its associated pointer type
https://stackoverflow.com/questions/11130592/golang-pointers-on-pointers-as-function-parameters/11131348#11131348
*/

func WaitGroup() {

	fmt.Println("OS \t", runtime.GOOS)
	fmt.Println("ARCH \t", runtime.GOARCH)
	fmt.Println("CPUs \t", runtime.NumCPU())
	fmt.Println("GoRoutines \t", runtime.NumGoroutine())
	fmt.Printf("%T\n", wg)

	wg.Add(1)
	// new goroutine
	go foo()
	bar()

	fmt.Println("uno")
	fmt.Println("dos")
	wg.Wait()
}

func foo() {
	fmt.Println("hello, I'm foo()")
	wg.Done()
}

func bar() {
	fmt.Println("I'm bar()")
}

func RaceCondition() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 10

	var wg sync.WaitGroup
	wg.Add(gs)

	var mutex sync.Mutex

	for i := 0; i < gs; i++ {
		fmt.Println("i: ", i)
		go func() {
			mutex.Lock()
			fmt.Println("counter", counter)
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mutex.Unlock()
			wg.Done()
		}()
		fmt.Println("counter", counter)
		fmt.Println("Go-routines: ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("count: ", counter)
}

func Exercise1() {
	fmt.Println("create 2 goroutines, and print something, use waitgroups")

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("I'm goroutine 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("I'm goroutine 2")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("finish")
}

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Println("hello, I'm person form type *person")
}

type human interface {
	speak()
}

func saySomething(h human) {
	fmt.Println("I'm inside saySomething func")
	h.speak()
}

func Exercise2() {
	p := person{
		name: "Mila",
		age:  22,
	}
	fmt.Println("type person cannot call saySomething")
	// saySomething(p) undefined error
	fmt.Println("type *person can call saySomething")
	saySomething(&p) // this works
}

func Exercise3() {

	var wg sync.WaitGroup
	wg.Add(10)
	var mutex sync.Mutex
	var n int = 0

	for i := 0; i < 10; i++ {
		go func() {
			mutex.Lock()
			local := n
			fmt.Println("original:", local)
			runtime.Gosched()
			local++
			n = local
			fmt.Println("incremented:", local)
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("finish:", n)
}

// using atomic
func Exercise4() {
	var wg sync.WaitGroup
	wg.Add(10)

	var n int64 = 0

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("original:", n)
			atomic.AddInt64(&n, 1)
			fmt.Println("incremented:", atomic.LoadInt64(&n))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("finish:", n)
}

func Exercise5() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
