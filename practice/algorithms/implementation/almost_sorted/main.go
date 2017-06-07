package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	s := readInput(os.Stdin)
	m, err := sortMethod(s)
	if err == nil {
		fmt.Printf("yes\n%s\n", m)
	} else {
		fmt.Println("no")
	}
}

func readInput(r io.Reader) []int {
	var n int
	fmt.Fscanf(r, "%d\n", &n)
	s := make([]int, n)
	for i := range s {
		fmt.Fscanf(r, "%d", &s[i])
	}
	return s
}

func swap(s []int, l, r int) {
	s[l], s[r] = s[r], s[l]
}

func reverse(s []int) {
	for i := 0; i <= (len(s)-1)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

type method struct {
	name string
	l, r int
}

func (m *method) String() string {
	if m == nil {
		return ""
	}
	return fmt.Sprintf("%s %d %d", m.name, m.l+1, m.r+1)
}

var (
	errUnableSort = errors.New("unable sort")
)

func sortMethod(s []int) (*method, error) {
	l, r := -1, -1
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			l = i
			break
		}
	}
	for i := len(s) - 1; i > 0; i-- {
		if s[i] < s[i-1] {
			r = i
			break
		}
	}
	if l == 0 && r == 0 {
		return nil, nil
	}
	if l > -1 && r > -1 {
		swap(s, l, r)
		if sort.IntsAreSorted(s) {
			return &method{"swap", l, r}, nil
		}
		swap(s, l, r)
		reverse(s[1 : r+1])
		if sort.IntsAreSorted(s) {
			return &method{"reverse", l, r}, nil
		}
	}
	return nil, errUnableSort
}
