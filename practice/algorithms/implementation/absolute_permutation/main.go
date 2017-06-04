package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	ps := readInput(os.Stdin)
	for _, p := range ps {
		if a, err := p.absolute3(); err != nil {
			fmt.Println(-1)
		} else {
			// buffer is useful
			var buf bytes.Buffer
			for i, v := range a {
				fmt.Fprint(&buf, v)
				if i < len(a)-1 {
					fmt.Fprint(&buf, " ")
				}
			}
			fmt.Printf("%s\n", &buf)
		}
	}
}

func readInput(r io.Reader) []*p {
	var n int
	fmt.Fscanf(r, "%d\n", &n)
	ps := make([]*p, n)
	for i := 0; i < n; i++ {
		ps[i] = &p{}
		fmt.Fscanf(r, "%d %d\n", &ps[i].n, &ps[i].k)
	}
	return ps
}

var errNoAbsolute = errors.New("no absolute permutation")

type p struct{ n, k int }

func (p *p) String() string {
	return fmt.Sprintf("[%d, %d]", p.n, p.k)
}

func (p *p) absolute3() ([]int, error) {
	lmap := make(map[int]map[int]struct{})
	oneChoice := make(map[int]struct{})
	for i := 1; i <= p.n; i++ {
		lmap[i] = make(map[int]struct{})
		pi0 := p.k + i
		pi1 := i - p.k
		if pi0 >= 1 && pi0 <= p.n {
			lmap[i][pi0] = struct{}{}
		}
		if pi1 >= 1 && pi1 <= p.n {
			lmap[i][pi1] = struct{}{}
		}
		if len(lmap[i]) == 0 {
			return nil, errNoAbsolute
		} else if len(lmap[i]) == 1 {
			oneChoice[i] = struct{}{}
		}
	}
	s := make([]int, p.n+1)
	umap := make(map[int]bool)
	fmap := make(map[int]bool)
	for len(oneChoice) > 0 {
		// fmt.Println(oneChoice)
		oneChoice2 := make(map[int]struct{})
		for k := range oneChoice {
			v := lmap[k]
			for v0 := range v {
				if umap[v0] {
					return nil, errNoAbsolute
				}
				s[k] = v0
				fmap[k] = true
				umap[v0] = true
				k0 := k + 2*p.k
				if k0 >= 1 && k0 <= p.n && !fmap[k0] {
					delete(lmap[k0], v0)
					oneChoice2[k0] = struct{}{}
				}
				k1 := k - 2*p.k
				if k1 >= 1 && k1 <= p.n && !fmap[k1] {
					delete(lmap[k1], v0)
					oneChoice2[k1] = struct{}{}
				}
			}
		}
		oneChoice = oneChoice2
	}
	if len(fmap) != p.n {
		return nil, errNoAbsolute
	}
	return s[1:], nil
}

func (p *p) absolute2() ([]int, error) {
	lmap := make(map[int]map[int]struct{})
	oneChoice := make(map[int]struct{})
	for i := 1; i <= p.n; i++ {
		lmap[i] = make(map[int]struct{})
		pi0 := p.k + i
		pi1 := i - p.k
		if pi0 >= 1 && pi0 <= p.n {
			lmap[i][pi0] = struct{}{}
		}
		if pi1 >= 1 && pi1 <= p.n {
			lmap[i][pi1] = struct{}{}
		}
		if len(lmap[i]) == 0 {
			return nil, errNoAbsolute
		} else if len(lmap[i]) == 1 {
			oneChoice[i] = struct{}{}
		}
	}
	s := make([]int, p.n+1)
	umap := make(map[int]bool)
	for len(oneChoice) > 0 {
		for k := range oneChoice {
			v := lmap[k]
			for v0 := range v {
				if umap[v0] {
					return nil, errNoAbsolute
				}
				s[k] = v0
				umap[v0] = true
			}
		}
		oneChoice = make(map[int]struct{})
		for i, l := range lmap {
			if len(l) == 0 {
				continue
			}
			l0 := make([]int, 0)
			for v0 := range l {
				l0 = append(l0, v0)
			}
			for _, v0 := range l0 {
				if umap[v0] {
					delete(l, v0)
				}
			}
			if len(l) == 1 {
				oneChoice[i] = struct{}{}
			}
		}
	}
	return s[1:], nil
}

func (p *p) absolute() ([]int, error) {
	lmap := make([][]int, p.n+1)
	for i := 1; i <= p.n; i++ {
		lmap[i] = make([]int, 0)
		pi0 := p.k + i
		pi1 := i - p.k
		if pi0 >= 1 && pi0 <= p.n {
			lmap[i] = append(lmap[i], pi0)
		}
		if pi1 >= 1 && pi1 <= p.n {
			lmap[i] = append(lmap[i], pi1)
		}
	}
	for i := 1; i <= p.n; i++ {
		if len(lmap[i]) == 0 {
			return nil, errNoAbsolute
		}
	}
	s := make([]int, p.n+1)
	umap := make(map[int]bool)
	if err := p.absolute0(s, 1, lmap, umap); err != nil {
		return nil, errNoAbsolute
	}
	return s[1:], nil
}

func (p *p) absolute0(s []int, i int, lmap [][]int, umap map[int]bool) error {
	if i > p.n {
		return nil
	}
	np := lmap[i]
	for _, v := range np {
		if umap[v] {
			continue
		}
		s[i] = v
		umap[v] = true
		err := p.absolute0(s, i+1, lmap, umap)
		if err != nil {
			s[i] = 0
			umap[v] = false
		} else {
			return nil
		}
	}
	return errNoAbsolute
}
