package main

import (
	"fmt"
	"myapp/pkg/handlers"
	"net/http"
)

func main() {
	fmt.Println("start :8080")
	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_ = http.ListenAndServe(":8080", nil)
}
