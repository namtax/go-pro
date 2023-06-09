package popcount

import "fmt"

var pc [256]byte

func init() {
	fmt.Println("------")
	// fmt.Println(pc)
	for i := range pc {
		// fmt.Println("-----")
		// fmt.Printf("i = %v\n", i)
		// fmt.Printf("bytes = %v\n", byte(i&1))
		// fmt.Println("-----")
		pc[i] = pc[i/2] + byte(i&1)
	}
	// fmt.Println(pc)
}

func PopCount(x uint64) int {
	fmt.Printf("x = %v\n", x)
	dd := byte(x >> (0 * 8))
	fmt.Printf("dd = %v\n", dd)
	fmt.Printf("current bytes = %v\n", byte(x>>(0*8)))
	fmt.Printf("current bytes = %v\n", byte(x>>(1*8)))
	fmt.Printf("current bytes = %v\n", byte(x>>(2*8)))
	fmt.Printf("current bytes = %v\n", byte(x>>(3*8)))
	fmt.Printf("current bytes = %v\n", byte(x>>(4*8)))
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
