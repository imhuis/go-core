package struct_type

import "learning-go/struct-type/abc"

type Point struct{ X, Y int }

var p = Point{1, 2}

var t = abc.T{
	// unable export
	//a: "",
	//b: "",
}
