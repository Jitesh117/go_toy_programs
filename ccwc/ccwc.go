package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var countWords bool
var countLines bool
var countChars bool
var countBytes bool

func init() {
	flag.BoolVar(&countWords, "w", false, "Count words")
	flag.BoolVar(&countLines, "l", false, "Count lines")
	flag.BoolVar(&countChars, "c", false, "Count characters")
	flag.BoolVar(&countBytes, "m", false, "Count bytes")
}

func main() {
	flag.Parse()
	args := flag.Args()

	if !countWords && !countLines && !countChars && !countBytes {
		countWords, countLines, countChars = true, true, true
	}

	if len(args) == 0 {
		count(os.Stdin, "")
	} else {
		for _, filename := range args {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "wc: %v\n", err)
				continue
			}
			count(file, filename)
			file.Close()
		}
	}
}

func count(r io.Reader, filename string) {
	scanner := bufio.NewScanner(r)
	var lines, words, chars, bytes int

	for scanner.Scan() {
		line := scanner.Text()
		lines++
		words += len(strings.Fields(line))
		chars += len(line)
		bytes += len([]byte(line)) + 1 // +1 for the newline character
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wc: %v\n", err)
	}

	output := ""
	if countLines {
		output += fmt.Sprintf("%d\t", lines)
	}
	if countWords {
		output += fmt.Sprintf("%d\t", words)
	}
	if countChars {
		output += fmt.Sprintf("%d\t", chars)
	}
	if countBytes {
		output += fmt.Sprintf("%d\t", bytes)
	}
	if filename != "" {
		output += filename
	}

	fmt.Println(output)
}
