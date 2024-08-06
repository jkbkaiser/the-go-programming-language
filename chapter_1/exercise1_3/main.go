package main

import (
    "fmt"
    "strings"
    "time"
)

func echo1(args []string) {
    var s, sep string
    for i := 1; i < len(args); i++ {
        s += sep + args[i]
        sep = " "
    }
    fmt.Println(s)
}

func echo2(args []string) {
    var s, sep string
    for _, arg := range args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}

func echo3(args []string) {
    fmt.Println(strings.Join(args[1:], " "))
}

func main() {
	i := 0
    num_iterations := 100000
    var args []string
    for i := 0; i < 100; i++ {
        args = append(args, "TEST")
    }

	start := time.Now()
	for i = 0; i < num_iterations; i++ {
		echo1(args)
	}
	resultEcho1 := time.Since(start).Seconds()

	start = time.Now()
	for i = 0; i < num_iterations; i++ {
		echo2(args)
	}
	resultEcho2 := time.Since(start).Seconds()

	start = time.Now()
	for i = 0; i < num_iterations; i++ {
		echo3(args)
	}
	resultEcho3 := time.Since(start).Seconds()

	fmt.Println(resultEcho1)
	fmt.Println(resultEcho2)
	fmt.Println(resultEcho3)
}
