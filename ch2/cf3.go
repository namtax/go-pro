package main

import (
	"fmt"

	"gopl.io/ch2/popcount"
)

func main() {
	setbits := popcount.PopCount(100)
	fmt.Printf("setbits = %v\n", setbits)
}
