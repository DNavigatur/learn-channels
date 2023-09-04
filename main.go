package main

import "fmt"

/*
// hands on exercise #1(a) :- bidirectional channels
func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}


// hands on exercise #1(b) :- buffer channels

func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}



// hands on exercise #2 :- buffer channels
func main() {
	cs := make(chan int)

	go func() {
		cs <- 42
		}()
		fmt.Println(<-cs)

		fmt.Printf("------\n")
		fmt.Printf("cs\t%T\n", cs)
	}

	// hands on exercise #3 :- range channels

	func main() {
		c := gen()
		receive(c)

		fmt.Println("about to exit")
	}

	func gen() <-chan int {
		c := make(chan int)

		go func() {
			for i := 0; i < 100; i++ {
				c <- i
		}
		close(c)
		}()

		return c
	}

	func receive(c <-chan int) {

		for v := range c {
			fmt.Println(v)
		}
	}

	// hands on exercise #4 :- select statement
	func main() {
		q := make(chan int)
		c := gen(q)

		receive(c, q)

		fmt.Println("about to exit")
	}

	func receive(c, q <-chan int) {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-q:
				return
			}
		}
	}

	func gen(q chan<- int) <-chan int {
		c := make(chan int)

		go func() {
			for i := 0; i < 100; i++ {
				c <- i
			}
			q <- 1
			close(c)
			}()

			return c
		}
		// hands on exercise #5 :- comma ok idiom
		func main() {
	c := make(chan int)

	go func() {
		c <- 42
		}()

		v, ok := <-c
		fmt.Println(v, ok)

		close(c)

		v, ok = <-c
		fmt.Println(v, ok)
	}

	// hands on exercise #6

	func main() {
		c := make(chan int)

		go func() {
			for i := 0; i < 100; i++ {
				c <- i
		}
		close(c)
		}()

		//recieving value from chan c
		for v := range c {
			fmt.Println(v)
		}
		fmt.Println("about to exit")

	}
*/

// hands on exercise #7

func main() {
	c := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			close(c)
		}()
	}
	for k := 0; k < 100; k++ {
		fmt.Println(<-c)
	}
	fmt.Println("about to exit")

}
