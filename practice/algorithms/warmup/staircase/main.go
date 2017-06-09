package main

import (
	"bytes"
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	fmt.Println(staircase(n))
}

func staircase(n int) string {
	var buf bytes.Buffer
	for i := 1; i <= n; i++ {
		fmt.Fprintf(&buf, fmt.Sprintf("%%%ds", n), multiple('#', i))
		if i < n {
			fmt.Fprintln(&buf)
		}
	}
	return buf.String()
}

func multiple(b byte, n int) string {
	bs := make([]byte, n)
	for i := range bs {
		bs[i] = b
	}
	return string(bs)
}
