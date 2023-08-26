package main

import "fmt"

type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string)
	fn := func(i int) {
		ch <- vs[i](name)
	}
	for i, _ := range vs {
		go fn(i)
	}
	return <-ch
}

func main1() {

	ret := exec("987654321", func(s string) string {
		return s + "func11"
	}, func(s string) string {
		return s + "func2"
	}, func(s string) string {
		return s + "func3"
	}, func(s string) string {
		return s + "func44"
	})
	fmt.Println(ret)

}
