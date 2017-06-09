package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

func main() {
	ss := readInput(os.Stdin)
	for _, s := range ss {
		bs, err := beautiful(s)
		if err == nil {
			fmt.Printf("YES %d\n", bs[0])
		} else {
			fmt.Println("NO")
		}
	}
}

func readInput(r io.Reader) []string {
	var n int
	fmt.Fscanf(r, "%d", &n)
	ss := make([]string, n)
	for i := range ss {
		fmt.Fscanf(r, "%s\n", &ss[i])
	}
	return ss
}

var errNotBeautiful = errors.New("string not beautiful")

func beautiful(s string) ([]int, error) {
	bs := []byte(s)
	for i := 1; i <= len(bs)/2; i++ {
		var ss []int
		j := 0
		curL := i
		for curL+j < len(bs) {
			c0, err := strconv.Atoi(string(bs[j : curL+j]))
			if err != nil {
				return nil, err
			}
			c1s := strconv.Itoa(c0 + 1)
			if curL+j+len(c1s) > len(bs) {
				break
			}
			if !reflect.DeepEqual([]byte(c1s), bs[curL+j:curL+j+len(c1s)]) {
				break
			} else {
				ss = append(ss, c0)
				j = j + curL
				curL = len(c1s)
				if j+curL == len(bs) {
					ss = append(ss, c0+1)
					return ss, nil
				} else if j+curL > len(bs) {
					break
				}
			}
		}
	}
	return nil, errNotBeautiful
}
