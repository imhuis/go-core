package struct_type

import (
	_ "encoding/json"
	"fmt"
	"time"
)

type Employee struct {
	ID         int
	Name       string
	Address    string
	Dob        time.Time
	Position   string
	Salary     float64
	PersonInfo Person `json1:"person"`
	Organization
	EmployID int64 `json1:"employ_id"`
}

type Organization struct {
	o string
}

type Person struct {
	ID        int    `json1:"id"`
	FirstName string `json1:"first_name"`
	LastName  string `json1:"last_name"`
	Address   string `json1:"address"`
}

type Employee1 struct {
	Person
	EmployeeNum int64
}

func main2() {
	fmt.Println("")
	var employee Employee = Employee{
		PersonInfo: Person{
			ID:        0,
			FirstName: "abc",
			LastName:  "cba",
			Address:   "abc",
		},
		EmployID: 0,
	}
	fmt.Println(employee.PersonInfo.ID)
}
