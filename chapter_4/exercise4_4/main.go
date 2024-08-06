package main

import "fmt"

func rotate(s []int, i int) {
    tmp := make([]int, i)
    copy(tmp, s[len(s) - i:])
	copy(s[i:], s[:len(s)-i])
    copy(s[:i], tmp)
}

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(a)
	rotate(a, 2)
	fmt.Println(a)
}
