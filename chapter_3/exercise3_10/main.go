package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	offset := len(s) % 3

	for i, v := range s {
		buf.WriteRune(v)

		if (i + 1) % 3 == offset && i != len(s) - 1 {
			buf.WriteRune(',')
		}
	}

	return buf.String()
}

func main() {
	var a = comma("1234123413")
	fmt.Println(a)
}
