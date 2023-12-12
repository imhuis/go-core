package functions

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main1() {
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, error := in.ReadRune()
		if error == io.EOF {
			break
		}
		if error != nil {
			fmt.Println(error)
		}
		fmt.Printf("%v", r)
	}
}
