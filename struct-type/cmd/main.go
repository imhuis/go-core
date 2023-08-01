package main

import (
	"fmt"
	struct_type "learning-go/struct-type"
	"time"
)

var m struct {
	Name string
	Age  int
}

var mm = struct {
	Name string
	Age  int
}{
	Name: "",
}

func main() {
	e := &struct_type.Employee{
		ID:         0,
		Name:       "",
		Address:    "",
		Dob:        time.Time{},
		Position:   "",
		Salary:     0,
		PersonInfo: struct_type.Person{},
		EmployID:   0,
	}
	p := &e.Position
	*p += "Senior"
	fmt.Printf("%T", p)
	e.Position = "Senior" + e.Position

	m := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
