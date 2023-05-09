package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float64
}

func (student *Student) Graduate() {
	student.Name = student.Name + " S.KOM"
}

func main() {
	student := Student{1, "alwani", 3.6}
	fmt.Println(student.Name)
	student.Graduate()
	fmt.Println(student.Name)

}
