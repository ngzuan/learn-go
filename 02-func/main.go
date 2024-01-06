package main

import "fmt"

func main() {
	x := sum(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("is total", x)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("index ", i, " total ", sum)
	}
	return sum

}
