package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	nums := make([]int, n)
	for i := 0; i < len(nums); i++ {
		fmt.Scanf("%d", &nums[i])
	}

	min := -1
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				if min == -1 || min > j-i {
					min = j - i
				}
				break
			}
		}
		if min == 1 {
			break
		}
	}
	fmt.Printf("%d", min)
}
