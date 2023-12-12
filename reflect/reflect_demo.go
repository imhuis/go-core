package reflect

import (
	"fmt"
	"reflect"
)

type Fruit struct {
	Name        string
	EnglishName string
}

type Apple struct {
	Fruit
}

func (f *Fruit) GetEnglishName() string {
	return f.EnglishName
}

//type IFruit interface {
//	getName() string-type
//}

func (f *Fruit) GetName() string {
	return f.Name
}

func printStructField(value reflect.Value) {
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("name: %v\n", value.Type().Field(i).Name)
		fmt.Printf("type: %v\n", value.Field(i).Type())
		fmt.Printf("type: %v\n", value.Type().Field(i).Type)
		fmt.Printf("value: %v\n", value.Field(i))
	}
}

func printStructMethod(value reflect.Value) {
	for i := 0; i < value.NumMethod(); i++ {
		fmt.Printf("name: %v\n", value.Method(i).String())
		fmt.Printf("name: %v\n", value.Type().Method(i).Name)

		fmt.Printf("type: %v\n", value.Method(i).Type())
		fmt.Printf("type: %v\n", value.Type().Method(i).Type)
	}
}
