package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"os"
	"strings"
)

const (
	baseURL     = "http://short.url/"
	shortLength = 8
)

func main() {
	var inputURL string

	if len(os.Args) > 1 {
		// Use command-line argument
		inputURL = os.Args[1]
	} else {
		// Read from standard input
		fmt.Print("Enter URL to shorten: ")
		reader := bufio.NewReader(os.Stdin)
		var err error
		inputURL, err = reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}
		inputURL = strings.TrimSpace(inputURL)
	}

	// Validate URL
	_, err := url.ParseRequestURI(inputURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid URL: %v\n", err)
		os.Exit(1)
	}

	// Generate short URL
	shortURL := generateShortURL(inputURL)

	fmt.Printf("Original URL: %s\n", inputURL)
	fmt.Printf("Shortened URL: %s\n", shortURL)
}

func generateShortURL(longURL string) string {
	// Create SHA-256 hash of the long URL
	hash := sha256.Sum256([]byte(longURL))

	// Encode the first 6 bytes of the hash to base64
	encoded := base64.URLEncoding.EncodeToString(hash[:6])

	// Use the first 8 characters of the encoded string
	shortCode := encoded[:shortLength]

	return baseURL + shortCode
}
