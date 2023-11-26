package struct_type

import "fmt"

func main3() {
	studentAge := map[string]int{
		"bob":  21,
		"alex": 20,
	}
	fmt.Println(studentAge)

	studentAge2 := make(map[string]int)
	fmt.Println(studentAge2)

	age, exist := studentAge["alex"]
	fmt.Println(age, exist)

}