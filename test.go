package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			close(c)
			return
		default:
			fmt.Println("default")
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {

		quit <- 0
		for i := range c {
			if i > 10 {
				break
			}
			fmt.Println(i)
		}

	}()
	fibonacci(c, quit)
}
