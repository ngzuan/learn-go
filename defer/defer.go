package main

import "fmt"

func main() {
	fmt.Println("defer đơn giản là tất cả các hàm return xong rồi mới thực hiện hàm đó")
	fmt.Println("EXAMPLE : ")
	a := "<3"
	defer fmt.Println("func viết đầu tiên nhưng chứa defer", a)
	fmt.Println("------------------------------------------")
	fmt.Println("func viết sau")
	defer fmt.Println("Hàm cuối cùng nhưng vẫn chứa defer !!!")
}
