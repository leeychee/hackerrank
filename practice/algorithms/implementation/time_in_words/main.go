package main

import "fmt"

func main() {
	var h, m int
	fmt.Scanf("%d\n%02d\n", &h, &m)

	fmt.Printf("%s", timeInWords(h, m))
}

var numstr []string = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
	"twenty",
}

const (
	TOP_CLOCK = -1
	PAST      = 0
	TO        = 1
)

var pasto []string = []string{
	"past",
	"to",
}

func timeInWords(h, m int) string {
	h, m, p := transfer(h, m)
	if p == TOP_CLOCK {
		return fmt.Sprintf("%s o' clock", numstr[h])
	} else {
		if m == 15 {
			return fmt.Sprintf("quarter %s %s", pasto[p], numstr[h])
		} else if m <= 20 {
			return fmt.Sprintf("%s minutes %s %s", numstr[m], pasto[p], numstr[h])
		} else if m < 30 {
			return fmt.Sprintf("%s %s minutes %s %s", numstr[m/10*10], numstr[m%10], pasto[p], numstr[h])
		} else {
			return fmt.Sprintf("half %s %s", pasto[p], numstr[h])
		}
	}
}

func transfer(h, m int) (int, int, int) {
	if m == 0 {
		return h, 0, TOP_CLOCK
	} else if m <= 30 {
		return h, m, PAST
	} else {
		return h + 1, 60 - m, TO
	}
}
