// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for filename, innerRange := range counts {
		for line, n := range innerRange {
			if n > 1 {
				fmt.Printf("%d\t%s\t%s", n, line, filename)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {

		if counts[f.Name()] == nil {
			counts[f.Name()] = make(map[string]int)
		}
		counts[f.Name()][input.Text()]++
	}
}
