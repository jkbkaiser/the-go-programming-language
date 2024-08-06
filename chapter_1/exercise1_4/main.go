package main

import (
    "bufio"
    "fmt"
    "os"
)

func countLines(f *os.File, occurences map[string]map[string]int, filename string) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        key := input.Text()
        if _, ok := occurences[key]; !ok {
            occurences[key] = make(map[string]int)
        }
        
        occurences[key][filename]++
    }
}

func main() {
    occurences := make(map[string]map[string]int)
    files := os.Args[1:]

    if len(files) == 0 {
        countLines(os.Stdin, occurences, "")
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)

            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }

            countLines(f, occurences, arg)
            f.Close()
        }
    }
    
    for line, entry := range occurences {
        var filenames []string
        var count int
        
        for filename, num := range entry {
            filenames = append(filenames, filename)
            count += num
        }

        if count > 1 {
            fmt.Printf("%d\t%v\t%s\n", count, filenames, line)
        }
    }
}
