package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("give one argument, an ASCII string of 2^N letters, 0<=N<=3")
	}
	str := os.Args[1]
	n := len(str)
	if !(n == 1 || n == 2 || n == 4 || n == 8) {
		panic("give one argument, an ASCII string of 2^N letters, 0<=N<=3")
	}
	var u uint64
	var shift uint
	for _, r := range str {
		if !(0 <= r && r <= 127) {
			panic("give one argument, an ASCII string of 2^N letters, 0<=N<=3")
		}
		u += uint64(r) << shift
		shift += 8
	}
	fmt.Print(u)
}
