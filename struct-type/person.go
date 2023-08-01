package struct_type

import "fmt"

func main1() {
	var person Person = Person{1, "Alex", "new", ""}
	fmt.Println(person)
	person = Person{ID: 1, FirstName: "old", LastName: ""}
	personCopy := &person
	personCopy.ID = 2
	fmt.Println(person.ID) //2

}
