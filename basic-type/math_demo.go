package basictype

import (
	"fmt"
	"math"
)

const (
	PI      = math.Pi
	IPv4Len = 4
)

func test_Nan() {
	nan := math.NaN()
	// false false false
	fmt.Println(nan == nan, nan > nan, nan < nan)
}

func ipv4(ips string) {
	var ip [IPv4Len]string
	fmt.Println(ip)
}
