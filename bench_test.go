package main

import (
	"fmt"
	_ "runtime/pprof"
	"testing"
)

func BenchmarkCrawler(t *testing.B) {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += i
	}
	fmt.Println(sum)
}

