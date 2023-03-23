/*
	Rewrite PopCount to use a loop instead of a single expression. Compare the performance of the two versions.
	(Section 11.4 shows how to compare the performance of different implementations systematically.)
*/

package popcount

import "fmt"

func init() {
	fmt.Println("------")

	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	// fmt.Println(pc)
}

func PopCountLoop(x uint64) byte {
	var resp byte

	for i := 0; i < 8; i++ {
		resp += pc[byte(x>>(i*8))]
	}

	return resp
}
