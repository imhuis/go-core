package funcs

import (
	"fmt"
	"strconv"
)

func summ(a, b string) (result int) {
	int1, _ := strconv.Atoi(a)
	int2, _ := strconv.Atoi(b)
	result = int1 + int2
	return
}

func main3() {
	var name *string
	*name = "adad"
	firstName := "steven"
	updateName(&firstName)
	fmt.Printf("First Name: %s", firstName)
}

func updateName(name *string) {
	*name = "yangz"
}
