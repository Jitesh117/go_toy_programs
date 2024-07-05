package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define the -n flag
	printLineNumbers := flag.Bool("n", false, "Print line numbers")
	flag.Parse()

	// Get the remaining arguments after the flags are parsed
	args := flag.Args()

	// Function to print lines from a file
	printLines := func(file *os.File) {
		scanner := bufio.NewScanner(file)
		for lineNumber := 1; scanner.Scan(); lineNumber++ {
			line := scanner.Text()
			if *printLineNumbers {
				fmt.Printf("%d\t%s\n", lineNumber, line)
			} else {
				fmt.Println(line)
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error: ", err)
		}
	}

	// If no file is specified, read from standard input
	if len(args) == 0 {
		printLines(os.Stdin)
	} else {
		// Iterate over each file argument
		for _, arg := range args {
			file, err := os.Open(arg)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error: ", err)
				continue
			}
			defer file.Close()
			printLines(file)
			fmt.Println()
		}
	}
}

