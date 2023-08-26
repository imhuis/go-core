package range_precautions

import "fmt"

type student struct {
	Name string
	Age  int
}

var (
	stus = []student{
		{
			Name: "zhou",
			Age:  20,
		},
		{
			Name: "yang",
			Age:  22,
		},
		{
			Name: "xh",
			Age:  23,
		},
	}
)

func listStu() {
	m := make(map[string]*student)
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	for _, stu := range m {
		fmt.Printf("%v", stu)
	}
}
