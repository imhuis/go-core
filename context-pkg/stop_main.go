package context_pkg

import (
	"fmt"
	"os"
)

func mainRoutine() {

	// 1.死循环
	//for{}

	// 2.select
	select {}

	// 3.
	c := make(chan struct{})
	c <- struct{}{}

	// 4.
	var cc chan struct{}
	<-cc

	// 5. 阻塞IO
	os.Stdin.Read(make([]byte, 1))

	// 6. 阻塞IO
	fmt.Scanln()
}
