package main

import (
	"container/heap"
	"fmt"
	"io"
	"os"
	"strings"
)

// HuffmanNode represents a node in the Huffman tree
type HuffmanNode struct {
	char  byte
	freq  int
	left  *HuffmanNode
	right *HuffmanNode
}

// PriorityQueue implements the heap.Interface for HuffmanNodes
type PriorityQueue []*HuffmanNode

// These methods are created to use the Heap container later
func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].freq < pq[j].freq
}
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*HuffmanNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// buildHuffmanTree builds the Huffman tree for the given text
func buildHuffmanTree(text string) *HuffmanNode {
	freqMap := make(map[byte]int)
	for i := 0; i < len(text); i++ {
		freqMap[text[i]]++
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	for char, freq := range freqMap {
		heap.Push(&pq, &HuffmanNode{char: char, freq: freq})
	}

	for pq.Len() > 1 {
		left := heap.Pop(&pq).(*HuffmanNode)
		right := heap.Pop(&pq).(*HuffmanNode)
		newNode := &HuffmanNode{
			freq:  left.freq + right.freq,
			left:  left,
			right: right,
		}
		heap.Push(&pq, newNode)
	}

	return heap.Pop(&pq).(*HuffmanNode)
}

// generateHuffmanCodes generates the Huffman codes for each character
func generateHuffmanCodes(root *HuffmanNode, code string, codes map[byte]string) {
	if root == nil {
		return
	}

	if root.left == nil && root.right == nil {
		codes[root.char] = code
		return
	}

	generateHuffmanCodes(root.left, code+"0", codes)
	generateHuffmanCodes(root.right, code+"1", codes)
}

// compress compresses the text using the Huffman codes
func compress(text string, codes map[byte]string) string {
	var compressed strings.Builder
	for i := 0; i < len(text); i++ {
		compressed.WriteString(codes[text[i]])
	}
	return compressed.String()
}

// decompress decompresses the compressed string using the Huffman tree
func decompress(compressed string, root *HuffmanNode) string {
	var decompressed strings.Builder
	current := root

	for _, bit := range compressed {
		if bit == '0' {
			current = current.left
		} else {
			current = current.right
		}

		if current.left == nil && current.right == nil {
			decompressed.WriteByte(current.char)
			current = root
		}
	}

	return decompressed.String()
}

func main() {
	var text string

	if len(os.Args) > 1 {
		filename := os.Args[1]
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}
		text = string(data)
	} else {
		fmt.Println("Enter text to compress (press Ctrl+D to end input):")
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Printf("Error reading stdin: %v\n", err)
			return
		}
		text = string(data)
	}

	root := buildHuffmanTree(text)

	codes := make(map[byte]string)
	generateHuffmanCodes(root, "", codes)

	compressed := compress(text, codes)

	if len(os.Args) > 1 {
		filename := os.Args[1]
		outputFilename := filename + "_compressed.txt"
		err := os.WriteFile(outputFilename, []byte(compressed), 0644)
		if err != nil {
			fmt.Printf("Error writing compressed file: %v\n", err)
			return
		}
		fmt.Printf("Compressed file saved as: %s\n", outputFilename)
	} else {
		fmt.Printf("Compressed text: %s\n", compressed)
	}
}
