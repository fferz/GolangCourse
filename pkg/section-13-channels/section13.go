package section13

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

/* 1. Channels intro */

// CHANNELS BLOCK
/*

	// this wont' work -> I need to put and to take a value inside the channel both at the same time

	c := make(chan int)

	c <- 22 // I add a value but at this moment there is nothing listenning => Channel Blocks!
					// I cannot put something in the channel until there is someone listenning

	fmt.Println(<-c)

*/

// -------------------------------------------------------------------------------------------------

// CHANNELS WON'T BLOCK - option 1
/*

c := make(chan int)

go func() {
	c <- 22
} ()

fmt.Println(<-c)

// Here I have 2 goroutines concurrently running.
// the value will be added to the channel when the println is ready to take it off from the channel
*/

// -------------------------------------------------------------------------------------------------

// CHANNELS WON'T BLOCK - option 2
/*

c := make(chan int, 1)  // buffered channel

c <- 22
c <- 33 // this will be blocked

fmt.Println(<-c)

// the buffered channel allows to keep 1 value, regardles someone is listenning or not.
// in this channel, I cannot add 2 values, only 1, otherwise the channel will get blocked.

*/

/* 2. Directional channels */

// Bidirectional

/*

c := make(chan int, 1)

c <- 22 // send

fmt.Println(<-c) // receive

*/

// Directional
/*

// send only channel
c := make(chan <- int, 2) // I can only send values to the channel
c <- 2
fmt.Println(<-c) // won´t work

// receive only channel
c := make(<- chan int, 2)
c <- 22 // won't work
fmt.Println(<-c)

*/

func ChannelAssign() {

	c := make(chan int)    // bidirectional
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	// specific to general - NO
	// c = cr
	// c = cs
	// (chan int)(cr)
	// (chan int)(cs)

	// specific to specifi - NO
	// cr = cs
	// cs = cr

	// general to specific
	cr = c
	cs = c
	n1 := (chan<- int)(c)
	n2 := (<-chan int)(c)

	fmt.Printf("%T\n", n1)
	fmt.Printf("%T\n", n2)

}

// using channels
/*
c := make(chan int)

		// send
		go func() {
			c <- 22
		}()

		// receive
		go func() {
			fmt.Println(<-c)
		}()

	fmt.Println("exit")
	// this sometimes works and sometimes doesn't.
	// it works when the receive goroutine is active when the send goroutine executes.
*/

func UsingChannels() {
	c := make(chan int)

	// send
	go func() {
		c <- 22
	}()

	// receive
	fmt.Println(<-c)

	fmt.Println("exit")
}

func Range() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // if we don't close, the listenner keeps waiting for more values => deadblock!
	}()

	for x := range c {
		fmt.Println(x)
	}

	fmt.Println("exit")
}

func SelectStatement() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(even, odd, quit)
	receive(even, odd, quit)

	fmt.Println("exit")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 11; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	//close(e)
	//close(o)
	q <- 0
}

func receive(e, o, q <-chan int) {

	for {
		select {
		case v := <-e:
			fmt.Println("even:", v)
		case v := <-o:
			fmt.Println("odd:", v)
		case v, ok := <-q:
			fmt.Println("quit:", v, ok)
			return
		}

	}
}

func CommaOk() {
	c := make(chan int)

	go func() {
		c <- 22
	}()
	v, ok := <-c
	fmt.Println(v, ok)
}

func FanIn() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	// send
	go send1(even, odd)
	go receive1(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("exit")
}

func send1(e, o chan<- int) {
	for i := 0; i < 11; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

func receive1(even, odd <-chan int, fanin chan<- int) {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v
		}
		// fanin <- <- even (the same as the for range)
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		// fanin <- <- odd (the same as the for range)
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}

func Fanin2() {
	c := fanIn(boring("joe"), boring("Ann"))

	for i := 0; i < 11; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("exit")
}

func boring(msj string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msj, i)
			time.Sleep(time.Duration(1000000))
		}

	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		c <- <-input2
	}()

	return c
}

func FanOut() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("exit")
}

func populate(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	// I could limit the number of goroutines, this way:
	// const goroutines = 10
	// for i:= 0; i < goroutines; i++
	for v := range c1 {
		wg.Add(1)
		go func(v2 int) {
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}

// CONTEXT
func ContextIntro() {

	println(" - context")
	ctx := context.Background()

	fmt.Println(ctx)
	fmt.Printf("type: %T\n", ctx)
	fmt.Println("error: ", ctx.Err())

	// context with cancel
	println(" - context with cancel")
	ctx1, cancel := context.WithCancel(context.Background())

	fmt.Println(ctx1)
	fmt.Printf("type: %T\n", ctx1)
	fmt.Println("error: ", ctx1.Err())

	println(" - context with cancel after cancel()")
	cancel()

	fmt.Println(ctx1)
	fmt.Printf("type: %T\n", ctx1)
	fmt.Println("error: ", ctx1.Err())

}

func ContextExample() {

	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1: ", ctx.Err())
	fmt.Println("num gortins 1: ", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2: ", ctx.Err())
	fmt.Println("num gortins 2: ", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	cancel()
}

func ContextExample2() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished

	for v := range gen(ctx) {
		fmt.Println(v)
		if v == 5 {
			break
		}
	}

}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // returning not to leak goroutine
			case dst <- n:
				n++
			}

		}
	}()

	return dst
}

func Exercise1() {
	/*
		c := make(chan int)

		c <- 22

		fmt.Println(<-c)
	*/
	// 1. func literal

	c := make(chan int)

	go func() {
		c <- 22
	}()

	fmt.Println(<-c)

	// 2. buffered channel
	d := make(chan int, 1)

	d <- 33

	fmt.Println(<-d)

}

func Exercise2() {
	/*
		c := make(chan int)

		go func() {

			for i := 0; i < 10; i++ {
				c <- i
			}
			close(c)
		}()

		for v := range c {
			fmt.Println(v)
		}

		fmt.Print("exit")

	*/

	c := newChannel()

	printChannel(c)

	fmt.Print("exit")

}

func newChannel() <-chan int {
	c := make(chan int)

	go func() {

		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func printChannel(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func Exercise3() {
	stop := make(chan int)
	canal := make(chan int)

	populateChannel(canal, stop)

	readChannel(canal, stop)

	fmt.Println("exit")
}

func readChannel(canal, stop <-chan int) {
	for {
		select {
		case v := <-canal:
			fmt.Println("canal:", v)
		case v, ok := <-stop:
			fmt.Println("canal:", v, ok)
			return
		}
	}
}

func populateChannel(canal, stop chan<- int) {

	go func() {
		for j := 0; j < 10; j++ {
			fmt.Println("valor:", j)
			canal <- j
		}
		stop <- 0
		close(canal)
		close(stop)
	}()

}

func Exercise4() {
	c := make(chan int)

	go func() {
		c <- 1
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}

func Exercise5() {
	c := make(chan int)

	go func() {
		for i := 0; i < 11; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}

func Exercise6() {
	c := make(chan int)

	// launch 10 goroutines
	for i := 0; i < 11; i++ {
		go func() {
			// add 3 numbers to a channel
			for j := 0; j < 3; j++ {
				c <- j
			}
		}()

	}
	for k := 0; k < 30; k++ {
		fmt.Println(k)
	}

}
