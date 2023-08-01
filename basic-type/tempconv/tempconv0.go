package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreesingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() {
	fmt.Printf("%g°C\n", c)
}

func (f Fahrenheit) String() {
	fmt.Printf("%g°F\n", f)
}

func main() {
	c := Fahrenheit(20.1)
	fmt.Println(FToC(c))
	c1 := CToF(FreesingC)
	c2 := CToF(BoilingC)
	fmt.Println(c1, c2)
}
