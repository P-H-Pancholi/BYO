package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/P-H-Pancholi/BYO/go/ccwc/scanner"
)

func main() {
	var input *os.File
	var filename string
	lineOnly := flag.Bool("l", false, "count lines only")
	wordOnly := flag.Bool("w", false, "count words only")
	byteOnly := flag.Bool("c", false, "count bytes only")
	charOnly := flag.Bool("m", false, "count characters only")

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		input = os.Stdin
		filename = ""
	} else {
		filename := args[0]
		input, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer input.Close()
	}

	stats, err := scanner.Scan(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	if *lineOnly {
		fmt.Printf("%d %s\n", stats.LineCount, filename)
	} else if *wordOnly {
		fmt.Printf("%d %s\n", stats.WordCount, filename)
	} else if *byteOnly {
		fmt.Printf("%d %s\n", stats.BytesCount, filename)
	} else if *charOnly {
		fmt.Printf("%d %s\n", stats.CharCount, filename)
	} else {
		fmt.Printf("Lines: %d, Words: %d, Bytes: %d, Chars: %d\n", stats.LineCount, stats.WordCount, stats.BytesCount, stats.CharCount)
	}
}
