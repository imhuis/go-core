package struct_type

import (
	"go-core/complex-type/struct-type/external"
)

type Point struct{ X, Y int }

var p = Point{1, 2}

var t = external.EXTStruct{
	// unable export
	//a: "",
	//b: "",
}
