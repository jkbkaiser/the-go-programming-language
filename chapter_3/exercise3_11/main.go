package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer
	decimal_index := strings.Index(s, ".")
	offset := decimal_index % 3

	for i, v := range s {
		buf.WriteRune(v)

		if (i+1)%3 == offset && i < decimal_index-1 {
			buf.WriteRune(',')
		}
	}

	return buf.String()
}

func main() {
	var a = comma("-12341234133.13241234")
	fmt.Println(a)
}
