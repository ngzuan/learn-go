package main

import (
	"fmt"
	"log"
)

func main() {
	var a string
	a = "Green"
	log.Println("my string is set to", a)
	changeUsingPointer(&a)
	log.Println("after func call my string is set to", a)

	/**Example*/
	fmt.Println("Example :")
	i, j := 90, 100
	var e *int = &i
	// đia chi cua i vd 23432x2220
	fmt.Println(e, *e)
	//1 dia chi cua i =23432x2220 ,2 là gia tri cua i=90
	*e = 20
	fmt.Println(e, *e, i)
	// (dia chi cua so 20,20,20)
	f := &j
	fmt.Println(f, *f)
	// dia chi chua j , gia tri cua j

	e = f
	fmt.Println(*e, *f)
	//gia tri cua  j(e,f la dia chi)

	*f = 69
	fmt.Println(*e, *f)
	// 69,69
}

/**pointer &->địa chỉ ổ nhớ của biến
** *a -> gia tri a la dia chi
 */

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}
