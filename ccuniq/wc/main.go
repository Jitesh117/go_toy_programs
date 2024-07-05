package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Get the command line arguments
	args := os.Args[1:]

	var input *os.File
	var output *os.File
	var err error

	switch len(args) {
	case 0:
		// No arguments, use standard input and output
		input = os.Stdin
		output = os.Stdout
	case 1:
		// One argument, use it as input file and standard output
		input, err = os.Open(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
			os.Exit(1)
		}
		defer input.Close()
		output = os.Stdout
	case 2:
		// Two arguments, use first as input file and second as output file
		input, err = os.Open(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
			os.Exit(1)
		}
		defer input.Close()
		output, err = os.Create(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer output.Close()
	default:
		fmt.Fprintln(os.Stderr, "Usage: uniq [input_file] [output_file]")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)
	defer writer.Flush()

	var prevLine string
	firstLine := true

	for scanner.Scan() {
		line := scanner.Text()
		if firstLine || line != prevLine {
			fmt.Fprintln(writer, line)
			prevLine = line
			firstLine = false
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}