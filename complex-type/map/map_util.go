package _map

import "fmt"

func fun1() {
	args := make(map[string]string)
	fmt.Println(args)
	delete(args, "")
}

// map不可以比
func equal(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, av := range a {
		if bv, ok := b[k]; !ok || bv != av {
			return false
		}
	}
	return true
}
