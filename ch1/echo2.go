// Echo prints its command line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	fmt.Println(os.Args)
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	// fmt.Println(arg)
	fmt.Println(s)
}
