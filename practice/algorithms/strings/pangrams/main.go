package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	if isPangram(string(s)) {
		fmt.Println("pangram")
	} else {
		fmt.Println("not pangram")
	}
}

func isPangram(s string) bool {
	m := make(map[byte]struct{})
	for _, b := range []byte(strings.ToLower(s)) {
		if b >= 'a' && b <= 'z' {
			m[b] = struct{}{}
		}
	}
	return len(m) == 26
}
