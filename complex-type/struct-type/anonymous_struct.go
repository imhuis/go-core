package struct_type

import "fmt"

func anonymousStruct() {
	var person struct {
		name string
		age  int
	}

	person.name = "steven"
	person.age = 18

	aa := struct {
		name string
		age  int
	}{
		name: "steven",
		age:  18,
	}
	fmt.Println(aa)
}
