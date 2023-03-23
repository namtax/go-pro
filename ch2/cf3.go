package main

import (
	"fmt"

	"gopl.io/ch2/popcount"
)

func main() {
	setbits := popcount.PopCount(1000)
	// setbitsloop := popcount.PopCountLoop(42)
	// setbitsbits := popcount.PopCountBits(42)
	fmt.Printf("setbits = %v\n", setbits)
	// fmt.Printf("setbitsloop = %v\n", setbitsloop)
	// fmt.Printf("setbitsbits = %v\n", setbitsbits)
}
