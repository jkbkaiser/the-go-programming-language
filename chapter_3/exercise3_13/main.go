package main

import "fmt"

const (
    _  = iota * 1000
    KB
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)

func main() {
    fmt.Println(KB)
}
