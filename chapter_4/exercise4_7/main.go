package main

import (
	"fmt"
	"unicode/utf8"
)

func swapSequence(b []byte, start1 int, start2 int, size int) {
	for k := 0; k < size; k++ {
		b[start1+k], b[start2+k] = b[start2+k], b[start1+k]
	}
}

// find middle, then copy
func reverse(b []byte) {
	var r1, r2 rune
	var s1, s2 int
    var roomFront, roomBack int
    var frontBuf, backBuf []rune

	i, j := 0, len(b)

	for i < j {
        if roomFront > 0 || roomBack > 0 {}

		r1, s1 = utf8.DecodeRune(b[i:])
		r2, s2 = utf8.DecodeLastRune(b[:j])

		if s1 == s2 {
            swapSequence(b, i, j-s2, s1)
            i += s1
            j -= s2
		} else if s1 > s2 {
            
            //backBuf = append(backBuf, r1)
            // do full swap
            copy(b[i:i+s2], b[j-s2:j])
            copy(b[i:i+s2], b[j-s2:j])
            
            i += s1
            roomFront = s1-s2
		} else {
            frontBuf = append(frontBuf, r2)
            copy(b[j-s1:j], b[i:i+s1])
            j += s1
            roomFront = s2-s2
        }
	}
}

func main() {
	//a := []byte("this is a test")
	//a := []byte("this is a 😀test")
	a := []byte("😀this is a b")
	fmt.Println(string(a))
	reverse(a)
	fmt.Println(string(a))
}
