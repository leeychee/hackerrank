package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(timeConv(s))
}

func timeConv(str string) string {
	var h, m, s int
	var ap string
	fmt.Sscanf(str, "%02d:%02d:%2d%1sM", &h, &m, &s, &ap)
	if ap == "P" {
		if h != 12 {
			h += 12
		}
	}
	if ap == "A" && h == 12 {
		h = 0
	}
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
