package main

//package main

func main() {

}

// 数组归零
func zero(ptr [32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func zero1(ptr *[3]int) {
	ptr = &[3]int{}
}
