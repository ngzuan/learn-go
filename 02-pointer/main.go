package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]User)

	me := User{
		FirstName: "Tuan",
		LastName:  "Nguyen",
	}
	myMap["name"] = me
	log.Println(myMap["name"].FirstName + myMap["name"].LastName)

}
