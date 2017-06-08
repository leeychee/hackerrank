package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scanln(&s)
	fmt.Println(changed(s))
}

func changed(s string) int {
	sos := []byte("SOS")
	bs := []byte(s)
	c := 0
	for i := 0; i < len(bs); i++ {
		if bs[i] != sos[i%3] {
			c++
		}
	}
	return c
}
