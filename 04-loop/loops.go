package main

import "fmt"

func main() {
	// loop bthuong
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	//lap ..
	// name := []string{"tuan", "nguyen", "van", "la", "hihi"}
	// for i, myname := range name {
	// 	fmt.Println(i, myname)
	// }

	// animal := make(map[string]string)
	// animal["dog"] = "cun"
	// animal["cat"] = "meo meo"
	// for i, myname := range animal {
	// 	fmt.Println(i, myname)
	// }
	// result dog cun
	// 		  cat meo meo

	var text string = "tuan dep trai"
	for i, t := range text {
		fmt.Println(i, t)
	}
}
