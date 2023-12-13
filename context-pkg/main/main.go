package main

import (
	"context"
	context_pkg "go-core/context-pkg"
	"os"
)

func main() {
	ss := context_pkg.SlowServer()
	defer ss.Close()
	fs := context_pkg.FastServer()
	defer fs.Close()

	ctx := context.Background()
	context_pkg.CallBoth(ctx, os.Args[1], ss.URL, fs.URL)
}
