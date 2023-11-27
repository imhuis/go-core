package rwmutex

import (
	"fmt"
	"sync"
	"time"
)

var muu sync.RWMutex
var count int

// panic fatal error: all goroutines are asleep - deadlock!
func main() {
	go A()
	time.Sleep(2 * time.Second)
	muu.Lock()
	defer muu.Unlock()
	count++
	fmt.Println(count)
}
func A() {
	muu.RLock()
	defer muu.RUnlock()
	B()
}
func B() {
	time.Sleep(5 * time.Second)
	C()
}
func C() {
	muu.RLock()
	defer muu.RUnlock()
}
