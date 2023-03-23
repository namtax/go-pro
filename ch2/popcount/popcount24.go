// Write a version of PopCount that counts bits by shifting its argument through 64 bit positions, testing the rightmost bit each time
package popcount

import "fmt"

func init() {
	fmt.Println("------")

	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	// fmt.Println(pc)
}

func PopCountBits(x uint64) int {
	var resp byte

	for i := 0; i < 64; i++ {
		current_byte := byte(x>>(i)) & 1
		fmt.Printf("current byte = %v\n", current_byte)
		resp += current_byte
	}

	return int(resp)
}
