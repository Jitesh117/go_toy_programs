package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func grep(pattern string, file *os.File, invertMatch bool) error {
	scanner := bufio.NewScanner(file)
	re, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	for scanner.Scan() {
		line := scanner.Text()
		match := re.MatchString(line)
		if (invertMatch && !match) || (!invertMatch && match) {
			fmt.Println(line)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func processFile(pattern string, filename string, invertMatch bool) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}
	defer file.Close()
	err = grep(pattern, file, invertMatch)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
}

func processDirectory(pattern string, dirname string, invertMatch bool) {
	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			processFile(pattern, path, invertMatch)
		}
		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
}

func main() {
	invertMatch := flag.Bool("v", false, "Invert match")
	recursive := flag.Bool("r", false, "Recursive search")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] PATTERN [FILE...]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	pattern := flag.Arg(0)
	if flag.NArg() == 1 {
		err := grep(pattern, os.Stdin, *invertMatch)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	} else {
		for _, filename := range flag.Args()[1:] {
			fileInfo, err := os.Stat(filename)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				continue
			}
			if fileInfo.IsDir() && *recursive {
				processDirectory(pattern, filename, *invertMatch)
			} else if fileInfo.IsDir() && !*recursive {
				fmt.Fprintf(os.Stderr, "Error: %s is a directory, use -r to search recursively\n", filename)
			} else {
				processFile(pattern, filename, *invertMatch)
			}
		}
	}
}
func nvimEditing() {
	fmt.Println("This code is being written in neovim lessgooo")

}
