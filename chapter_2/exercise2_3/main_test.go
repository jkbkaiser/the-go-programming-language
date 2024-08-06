package main 

import "testing"

func BenchmarkPopCountTable(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PopCountTable(uint64(i))
    }
}

func BenchmarkPopCountLoop(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PopCountTable2(uint64(i))
    }
}
