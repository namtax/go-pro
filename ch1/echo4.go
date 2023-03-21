package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newlines")
var sep = flag.String("s", " ", "seperator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	// fmt.Println(os.Args[0])
	// fmt.Println(os.Args)
	// fmt.Println(strings.Join(os.Args[1:], " "))

	// s, sp := "", ""
	// for idx, rng := range os.Args[1:] {
	// 	s += sp + strconv.Itoa(idx) + rng
	// }

	// println(s)
}
