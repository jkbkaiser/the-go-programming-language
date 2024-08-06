package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main() {
    start := time.Now()
    ch := make(chan string)

    for _, url :=  range os.Args[1:] {
        go fetch(url, ch)
    }

    f, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    defer f.Close()
    
    if err != nil {
        fmt.Fprintf(os.Stderr, "while opening file test.log: %s\n", err)
    }

    for range os.Args[1:] {
        if _, err := f.WriteString(<-ch); err != nil {
            fmt.Fprintf(os.Stderr, "while writing to file test.log: %s\n", err)
        }
    }

    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    
    if err != nil {
        ch <- fmt.Sprint(err)
        return
    }

    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close()

    if err != nil {
        ch <- fmt.Sprintf("while reading: %s: %v\n", url, err)
    }


    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d  %s\n", secs, nbytes, url)
}
