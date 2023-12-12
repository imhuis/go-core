package string_type

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	l := len(s)
	fmt.Println(l, s[l-1])

	fmt.Printf("实际码点数：%d\n", utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\t%d\n", i, r, r)
		i += size
	}

}
