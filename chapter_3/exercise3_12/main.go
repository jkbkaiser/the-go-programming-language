package main

import "fmt"

func isAnagram(a string, b string) bool {
    m := make(map[rune]int)

    for _, v := range a {
        m[v]++
    }

    for _, v := range b {
        m[v]--
    }
    
    for _, val := range m {
        if val != 0 {
            return false
        }
    }

    return true
}

func main() {
    fmt.Println(isAnagram("😀ba", "a😀b"))
}
