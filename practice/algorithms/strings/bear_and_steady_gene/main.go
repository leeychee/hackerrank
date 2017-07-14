package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	var n int
	fmt.Scanf("%d\n", &n)
	var s string
	fmt.Scanf(fmt.Sprintf("%%%ds\n", n), &s)

	_, subs := makeSteady1(s)
	fmt.Println(subs)
}

func steady(s string) bool {
	return false
}

func makeSteady(s string) (idx int, lgh int) {
	bs := []byte(s)
	n := len(s) / 4
	var na, nc, nt, ng int
	for _, b := range bs {
		switch b {
		case 'A':
			na++
		case 'C':
			nc++
		case 'T':
			nt++
		case 'G':
			ng++
		}
	}
	d := make(map[byte]int)
	da := na - n
	dc := nc - n
	dt := nt - n
	dg := ng - n
	d['A'] = da
	d['C'] = dc
	d['T'] = dt
	d['G'] = dg
	minLen := 0
	for _, v := range d {
		if v > 0 {
			minLen += v
		}
	}
	if minLen == 0 {
		return -1, 0
	}
	lm := make([][]int, len(bs))
	lgh = len(bs)
	pc := 0
	for i := 0; i < len(bs)-1; i++ {
		if d[bs[i]] > 0 && lm[i] == nil {
			lm[i] = make([]int, 4)
		}
		switch bs[i] {
		case 'A':
			for j := pc; j <= i && j < len(bs)-minLen; j++ {
				if i-j+1 >= lgh {
					pc = j + 1
					continue
				}
				m := lm[j]
				if m == nil {
					continue
				}
				m[0]++
				if m[0] >= da && m[1] >= dc && m[2] >= dt && m[3] >= dg {
					pc = j + 1
					l := i - j + 1
					if lgh > l {
						idx = j
						lgh = l
					}
				}
			}
		case 'C':
			for j := pc; j <= i && j < len(bs)-minLen; j++ {
				if i-j+1 >= lgh {
					pc = j + 1
					continue
				}
				m := lm[j]
				if m == nil {
					continue
				}
				m[1]++
				if m[0] >= da && m[1] >= dc && m[2] >= dt && m[3] >= dg {
					pc = j + 1
					l := i - j + 1
					if lgh > l {
						idx = j
						lgh = l
					}
				}
			}
		case 'T':
			for j := pc; j <= i && j < len(bs)-minLen; j++ {
				if i-j+1 >= lgh {
					pc = j + 1
					continue
				}
				m := lm[j]
				if m == nil {
					continue
				}
				m[2]++
				if m[0] >= da && m[1] >= dc && m[2] >= dt && m[3] >= dg {
					pc = j + 1
					l := i - j + 1
					if lgh > l {
						idx = j
						lgh = l
					}
				}
			}
		case 'G':
			for j := pc; j <= i && j < len(bs)-minLen; j++ {
				if i-j+1 >= lgh {
					pc = j + 1
					continue
				}
				m := lm[j]
				if m == nil {
					continue
				}
				m[3]++
				if m[0] >= da && m[1] >= dc && m[2] >= dt && m[3] >= dg {
					pc = j + 1
					l := i - j + 1
					if lgh > l {
						idx = j
						lgh = l
					}
				}
			}
		}
	}
	return
}

func makeSteady1(s string) (idx int, lgh int) {
	bs := []byte(s)
	n := len(s) / 4
	var na, nc, nt, ng int
	for _, b := range bs {
		switch b {
		case 'A':
			na++
		case 'C':
			nc++
		case 'T':
			nt++
		case 'G':
			ng++
		}
	}
	d := make(map[byte]int)
	da := na - n
	dc := nc - n
	dt := nt - n
	dg := ng - n
	d['A'] = da
	d['C'] = dc
	d['T'] = dt
	d['G'] = dg
	minLen := 0
	for _, v := range d {
		if v > 0 {
			minLen += v
		}
	}
	if minLen == 0 {
		return -1, 0
	}
	lgh = len(bs)
	for i := 0; i < len(bs)-minLen-1; i++ {
		if d[bs[i]] < 1 {
			continue
		}
		var sa, sc, st, sg int
		for j := i; j < len(bs) && j-i+1 < lgh; j++ {
			switch bs[j] {
			case 'A':
				sa++
			case 'C':
				sc++
			case 'T':
				st++
			case 'G':
				sg++
			}
			if sa >= da && sc >= dc && st >= dt && sg >= dg {
				if lgh > j-i+1 {
					idx = i
					lgh = j - i + 1
				}
				break
			}
		}
	}
	return
}

func makeSteady2(s string) (idx, lgh int) {
	return
}
