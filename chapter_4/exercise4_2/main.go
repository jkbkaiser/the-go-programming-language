package main

import (
	"crypto/sha256"
    "crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

var aFlag = flag.String("a", "256", "algorithm to use for hash")

func alg256(data []byte) []byte {
    res := sha256.Sum256(data)
    return res[:] 
}

func alg384(data []byte) []byte {
    res := sha512.Sum384(data)
    return res[:]
}

func alg224(data []byte) []byte {
    res := sha256.Sum224(data)
    return res[:]
}

func getAlg(flag string) func([]byte) []byte {
    if flag == "256" {
        return alg256
    }

    if flag == "224" {
        return alg384
    }

    if flag == "384" {
        return alg224
    }

    return nil
}

func main() {
	flag.Parse()
    alg := getAlg(*aFlag)

    input, err := io.ReadAll(os.Stdin)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading input")
    }
    res := alg(input)
    fmt.Println(input)
    fmt.Println(res)
}
