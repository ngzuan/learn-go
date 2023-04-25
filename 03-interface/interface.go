package main

import "log"

type Person interface {
	Say() string
	NumberOfAge() int
}
type Student struct {
	Name string
	Age  int
}
type Doctor struct {
	Name string
	Age  int
	job  string
}

func main() {
	student := Student{
		Name: "tuan",
		Age:  18,
	}
	PrintInfo(&student)
	doctor := Doctor{
		Name: "hihi",
		Age:  29,
		job:  "doctor",
	}
	PrintInfo(&doctor)

}

func PrintInfo(p Person) {
	log.Println("this is Person say", p.Say(), "person is", p.NumberOfAge(), "age")
}

func (s *Student) Say() string {
	return "Me"
}
func (s *Student) NumberOfAge() int {
	return 9
}

func (s *Doctor) Say() string {
	return "Doctor"
}
func (s *Doctor) NumberOfAge() int {
	return 88
}
