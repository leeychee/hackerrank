package main

import "fmt"

func main() {
	var n, d int
	fmt.Scanf("%d %d", &n, &d)

	a := make([]int, n)
	for i := 0; i < len(a); i++ {
		fmt.Scanf("%d", &a[i])
	}

	var r int
	for i := 0; i < len(a); i++ {
	loop:
		for j := i + 1; j < len(a); j++ {
			if a[j]-a[i] == d {
				for k := j + 1; k < len(a); k++ {
					if a[k]-a[j] == d {
						r++
					} else if a[k]-a[j] > d {
						break loop
					}
				}
			} else if a[j]-a[i] > d {
				break loop
			}
		}
	}

	fmt.Printf("%d", r)
}
