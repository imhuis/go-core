package type_aliase

import "time"

type T3 struct{}
type T4 = T3

func (t T3) say() {

}

type S struct {
	T3
	T4
}

type Time = time.Time

func main() {
	var s S
	s.say()
}
