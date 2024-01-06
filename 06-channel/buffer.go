package main

import "fmt"

func main() {
	b := make(chan string, 1) // co 1 stack
	b <- "tuan"               // chuyen vao stack
	fmt.Println(<-b)          // lay b ra -> stack rong
	b <- "nguyen"             // tiep tuc
	fmt.Println(<-b)          // stack rong

}
