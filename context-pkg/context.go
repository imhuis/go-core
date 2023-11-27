package main

import (
	"context"
	"os"
	"os/signal"
)

func main() {
	// 1. 捕获 os.Signal
	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt)
	//<-c

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	<-ctx.Done()
}
