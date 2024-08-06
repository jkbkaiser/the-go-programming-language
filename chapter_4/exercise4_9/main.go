package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	word_count := make(map[string]int)

	for in.Scan() {
		word := in.Text()
		word_count[word]++
	}

	for key, val := range word_count {
		fmt.Printf("word: %v\t count: %d\n", key, val)
	}
}
