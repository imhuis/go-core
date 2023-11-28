package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func Example_reflectBasicType() {
	t := reflect.TypeOf(3)
	fmt.Printf("%T, %v\n", t, t.String())
	// output:
	//*reflect.rtype, int
}

func Example_reflect() {
	f := Fruit{Name: "Apple"}

	valueType := reflect.ValueOf(f)
	pointType := reflect.ValueOf(&f)
	fmt.Printf("Type: %v, %v\n", valueType.Type(), valueType.String())
	fmt.Printf("point Type: %v\n", pointType.Type())

	for i := 0; i < valueType.NumField(); i++ {
		fmt.Printf("name:%v - type:%v\n", valueType.Type().Field(i).Name, valueType.Field(i).Type())
	}

	for i := 0; i < pointType.NumField(); i++ {
		fmt.Printf("name:%v - type:%v\n", pointType.Type().Field(i).Name, pointType.Field(i).Type())
	}

	// output:
	//Type: reflect.Fruit, <reflect.Fruit Value>
	//point Type: *reflect.Fruit
	//name:Name - type:string
	//name:EnglishName - type:string
}

func Test_struct(t *testing.T) {
	var f Fruit
	printStructField(reflect.ValueOf(f))
}

func Test_struct2(t *testing.T) {
	var f = Fruit{
		Name:        "",
		EnglishName: "",
	}
	printStructMethod(reflect.ValueOf(f))
}
