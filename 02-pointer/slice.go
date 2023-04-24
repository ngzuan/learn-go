package main

import (
	"log"
	"sort"
)

func main() {
	myAr := []int{1, 2, 3, 4, 6, 7, 0} // tao mang
	// lay phan tu trong mang
	log.Println(myAr[3:5])

	var mySlice []int
	mySlice = append(mySlice, 3)
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 4)
	mySlice = append(mySlice, 5)

	sort.Ints(mySlice)

	log.Println(mySlice)

}
