package main

import "fmt"

func foo() chan int {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()

	return ch
}

func main() {
	for n := range foo() {
		fmt.Println(n)
	}
	fmt.Println("channel is now closed")
}
