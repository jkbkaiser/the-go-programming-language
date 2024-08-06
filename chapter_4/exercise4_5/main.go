package main

import "fmt"

func removeDuplicates(s []string) []string {
    i := 0
    for _, val := range s {
        if val == s[i] {
            continue
        }
        i++
        s[i] = val
    }

    return s[:i + 1]
}

func main() {
    a := []string{"1", "b", "b", "b", "c"}
    fmt.Println(a)
    a = removeDuplicates(a)
    fmt.Println(a)
}
