// Modified dup2 for ex1.4
// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type dup struct {
	count int
	fname map[string]struct{}
}

func main() {
	dups := make(map[string]*dup)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, -1, dups)
	} else {
		for i, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, i, dups)
			f.Close()
		}
	}
	for line, val := range dups {
		n := val.count
		keys := []string{}
		for k := range val.fname {
			keys = append(keys, k)
		}
		if n > 1 {
			fmt.Printf("%v\t%v\t%s\n", val.count, keys, line)
		}
	}
}

func countLines(f *os.File, findex int, dups map[string]*dup) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		inp := input.Text()
		if val, exists := dups[inp]; exists {
			val.count++
			val.fname[os.Args[findex+1]] = struct{}{}
		} else {
			m := map[string]struct{}{}
			m[os.Args[findex+1]] = struct{}{}
			dups[inp] = &dup{1, m}
		}

	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
