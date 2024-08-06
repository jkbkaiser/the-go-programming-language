package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	letter_counts := make(map[rune]int)
	digit_counts := make(map[rune]int)
	punct_counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsLetter(r) {
			letter_counts[r]++
		}

		if unicode.IsDigit(r) {
			digit_counts[r]++
		}

		if unicode.IsPunct(r) {
			punct_counts[r]++
		}

		utflen[n]++
	}

	fmt.Printf("letter\tcount\n")
	for c, n := range letter_counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("digit\tcount\n")
	for c, n := range digit_counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("punctuation\tcount\n")
	for c, n := range punct_counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%dn\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 character\n", invalid)
	}
}
