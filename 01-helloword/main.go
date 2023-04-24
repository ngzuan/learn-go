package main

import "log"

func main() {
	var a string
	a = "Green"
	log.Println("my string is set to", a)
	changeUsingPointer(&a)
	log.Println("after func call my string is set to", a)

}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue

}
