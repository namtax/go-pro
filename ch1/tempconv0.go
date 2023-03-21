package main

import (
	"fmt"
	"reflect"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	// fmt.Printf("%g\n", BoilingC-FreezingC)
	// BoilingF := CToF(BoilingC)
	// fmt.Printf("%T", BoilingF)
	// fmt.Printf("%g\n", BoilingF-CToF(FreezingC))

	var c Celsius = 99
	// var f Fahrenheit = 0

	fmt.Println("-----")
	fmt.Println(c.String())
	fmt.Println(c.Type())
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToc(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string     { return fmt.Sprintf("%gC", c) }
func (c Celsius) Type() reflect.Type { return reflect.TypeOf(c) }
