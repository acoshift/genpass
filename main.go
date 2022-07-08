package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"strconv"
)

const char = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func main() {
	var n int
	if len(os.Args) > 1 {
		n, _ = strconv.Atoi(os.Args[1])
	}
	if n <= 0 {
		n = 32
	}

	var s string
	for i := 0; i < n; i++ {
		b := make([]byte, 1)
		rand.Read(b)
		s += string(char[int(b[0])%len(char)])
	}
	fmt.Print(s)
}
