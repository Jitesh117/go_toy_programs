package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

// readLines reads lines from the provided file or standard input
func readLines(file *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// unique removes duplicate lines from the slice
func unique(lines []string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, line := range lines {
		if !seen[line] {
			seen[line] = true
			result = append(result, line)
		}
	}
	return result
}

func main() {
	// Define flags
	uniqueFlag := flag.Bool("u", false, "print only unique lines")
	randomFlag := flag.Bool("r", false, "sort lines in random order")
	flag.Parse()

	var file *os.File
	var err error

	// Determine the source of the input (file or standard input)
	if len(flag.Args()) > 0 {
		file, err = os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: could not open file %s: %v\n", flag.Arg(0), err)
			return
		}
		defer file.Close()
	} else {
		file = os.Stdin
	}

	// Read lines from the file or standard input
	lines, err := readLines(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: reading input: %v\n", err)
		return
	}

	// Apply unique filter if -u flag is set
	if *uniqueFlag {
		lines = unique(lines)
	}

	// Sort lines randomly if -r flag is set, otherwise sort alphabetically
	if *randomFlag {
		randSource := rand.NewSource(time.Now().UnixNano())
		randGen := rand.New(randSource)
		randGen.Shuffle(len(lines), func(i, j int) { lines[i], lines[j] = lines[j], lines[i] })
	} else {
		sort.Strings(lines)
	}

	// Print the sorted (and possibly unique) lines
	for _, line := range lines {
		fmt.Println(line)
	}
}
