package main

import iopkg "go-core/io-pkg"

func main() {
	err := iopkg.ReadRune()
	if err != nil {
		return
	}
}
