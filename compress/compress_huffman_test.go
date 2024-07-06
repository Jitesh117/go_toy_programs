package main

import (
	"os"
	"testing"
)

func TestBuildHuffmanTree(t *testing.T) {
	text := "this is an example for huffman encoding"
	root := buildHuffmanTree(text)

	if root == nil {
		t.Error("Expected a non-nil root for the Huffman tree")
	}
}

func TestGenerateHuffmanCodes(t *testing.T) {
	text := "this is an example for huffman encoding"
	root := buildHuffmanTree(text)

	codes := make(map[byte]string)
	generateHuffmanCodes(root, "", codes)

	if len(codes) == 0 {
		t.Error("Expected non-zero length for Huffman codes")
	}

	expectedChars := []byte("this anexampleforufncoding")
	for _, char := range expectedChars {
		if _, exists := codes[char]; !exists {
			t.Errorf("Expected Huffman code for character '%c'", char)
		}
	}
}

func TestCompressAndDecompress(t *testing.T) {
	text := "this is an example for huffman encoding"
	root := buildHuffmanTree(text)

	codes := make(map[byte]string)
	generateHuffmanCodes(root, "", codes)

	compressed := compress(text, codes)
	if len(compressed) == 0 {
		t.Error("Expected non-zero length for compressed text")
	}

	decompressed := decompress(compressed, root)
	if decompressed != text {
		t.Errorf("Expected decompressed text to be '%s', but got '%s'", text, decompressed)
	}
}

func TestCompressFromFile(t *testing.T) {
	filename := "testfile.txt"
	originalText := "this is an example for huffman encoding"
	err := os.WriteFile(filename, []byte(originalText), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}
	defer os.Remove(filename)

	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}
	text := string(data)

	root := buildHuffmanTree(text)
	codes := make(map[byte]string)
	generateHuffmanCodes(root, "", codes)

	compressed := compress(text, codes)
	outputFilename := filename + "_compressed.txt"
	err = os.WriteFile(outputFilename, []byte(compressed), 0644)
	if err != nil {
		t.Fatalf("Failed to write compressed file: %v", err)
	}
	defer os.Remove(outputFilename)

	compressedData, err := os.ReadFile(outputFilename)
	if err != nil {
		t.Fatalf("Failed to read compressed file: %v", err)
	}

	decompressed := decompress(string(compressedData), root)
	if decompressed != originalText {
		t.Errorf("Expected decompressed text to be '%s', but got '%s'", originalText, decompressed)
	}
}
