package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte
const HashSize = 32

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func hashDiff(c1 [HashSize]byte, c2 [HashSize]byte) int {
    count := 0
    
    for i := 0; i < HashSize; i++ {
        count += int(pc[c1[i] ^ c2[i]])
    }

    return count;
}

func main() {
    c1 := sha256.Sum256([]byte("X"))
    c2 := sha256.Sum256([]byte("x"))
    fmt.Println(hashDiff(c1, c2))
    fmt.Println(c1)
    fmt.Println(c2)

    for _, val := range pc {
        fmt.Printf("%b\n", val)
    }
}
