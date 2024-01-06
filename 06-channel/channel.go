package main

import "fmt"

func main() {
	a := make(chan int)

	go func() {
		a <- 1
	}() // tao 1 goRoutime

	fmt.Println(<-a)
}
