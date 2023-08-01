package ioio

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadRune() error {
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("read failed: %v", err)
		}
		fmt.Printf("%c\n", r)
	}
	return nil
}
