package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(os.Args[0])
	fmt.Println(os.Args)
	fmt.Println(strings.Join(os.Args[1:], " "))

	s, sp := "", ""
	for idx, rng := range os.Args[1:] {
		s += sp + strconv.Itoa(idx) + rng
	}

	println(s)
}
