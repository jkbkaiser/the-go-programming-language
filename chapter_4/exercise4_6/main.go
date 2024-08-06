package main

import (
    "fmt"
    "unicode"
)

func squashSpace(b []byte) []byte {
    runes := []rune(string(b))
    i := 0
    for _, rune := range runes[1:] {
        if unicode.IsSpace(rune) && unicode.IsSpace(runes[i]) {
            continue
        }
        i++
        runes[i] = rune        
    } 
    return []byte(string(runes[:i+1]))
}

func main() {
    a := []byte("this is a  test")
    fmt.Println(string(a))
    a = squashSpace(a)
    fmt.Println(string(a))
}
