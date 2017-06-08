package main

import "fmt"

func main() {
	var s string
	fmt.Scanln(&s)
	r := superReduce(s)
	if r == "" {
		fmt.Println("Empty String")
	}
	fmt.Println(r)
}

func superReduce(s string) string {
	bs := []byte(s)
	bs = reduce(bs)
	return string(bs)
}

func reduce(bs []byte) []byte {
	reduced := true
	for reduced {
		reduced = false
		for i := 0; i < len(bs)-1; {
			if bs[i] == bs[i+1] {
				bs = append(bs[0:i], bs[i+2:]...)
				reduced = true
			} else {
				i++
			}
		}
	}
	return bs
}
