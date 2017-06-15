package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"sort"
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
	genes := readGenes(os.Stdin)
	cases := readCases(os.Stdin)
	min, max := minMaxHealth(genes, cases)
	fmt.Printf("%d %d\n", min, max)
}

func minMaxHealth(genes []gene, cases []_case) (int, int) {
	min := -1
	max := 0
	hub := newHub(genes)
	for _, c := range cases {
		h := hub.sub(c.start, c.end).health(c.dna)
		if min == -1 || h < min {
			min = h
		}
		if h > max {
			max = h
		}
	}
	return min, max
}

func readGenes(r io.Reader) []gene {
	var n int
	fmt.Fscanf(r, "%d\n", &n)
	genes := make([]gene, n)
	for i := 0; i < n; i++ {
		genes[i].idx = i
		fmt.Fscanf(r, "%s", &genes[i].dna)
	}
	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%d", &genes[i].val)
	}
	return genes
}

func readCases(r io.Reader) []_case {
	var t int
	fmt.Fscanf(r, "%d\n", &t)
	cases := make([]_case, t)
	for i := 0; i < t; i++ {
		var d _case
		fmt.Fscanf(r, "%d %d %s", &d.start, &d.end, &d.dna)
		cases[i] = d
	}
	return cases
}

type _case struct {
	start, end int
	dna        dna
}

type gene struct {
	idx int
	dna dna
	val int
}

type pGene struct {
	dna   dna
	pairs []pair
}

func (p *pGene) add(p0 pair) {
	p.pairs = append(p.pairs, p0)
}

type pGenes []pGene

func (p pGenes) Add(g gene) pGenes {
	if p == nil {
		p = pGenes{}
	}
	for i := 0; i < len(p); i++ {
		if p[i].dna == g.dna {
			p[i].add(pair{g.idx, g.val})
			return p
		}
	}
	p = append(p, pGene{g.dna, []pair{{g.idx, g.val}}})
	sort.Sort(p)
	return p
}

func (p pGenes) Less(i, j int) bool {
	return p[i].dna < p[j].dna
}

func (p pGenes) Len() int {
	return len(p)
}

func (p pGenes) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type mGenes struct {
	sm map[dna]*pGene
	lm map[int]struct{}
	ls []int
}

func (m *mGenes) add(g gene) *mGenes {
	if m == nil {
		m = &mGenes{}
	}
	if m.sm == nil {
		m.sm = make(map[dna]*pGene)
	}
	if _, ok := m.sm[g.dna]; !ok {
		m.sm[g.dna] = &pGene{g.dna, []pair{{g.idx, g.val}}}
	} else {
		m.sm[g.dna].add(pair{g.idx, g.val})
	}
	if m.lm == nil {
		m.lm = make(map[int]struct{})
	}
	m.lm[len(g.dna)] = struct{}{}
	return m
}

func (m *mGenes) lens() []int {
	if m.ls == nil {
		for l := range m.lm {
			m.ls = append(m.ls, l)
		}
	}
	return m.ls
}

func (m *mGenes) get(k dna) *pGene {
	return m.sm[k]
}

type pair struct{ idx, val int }

type geneHub struct {
	genes      []gene
	start, end int
	// match 1st byte, store the genes
	m1 [26]*mGenes
	// match 1st, 2nd bytes
	m2 [26][26]*mGenes
	// match 1st, 2nd, 3rd bytes
	m3 [26][26][26]*mGenes
	// real func
	health func(d dna) int
}

func newHub(gs []gene) *geneHub {
	h := &geneHub{
		genes: gs,
		end:   len(gs) - 1,
	}
	h.init()
	if maxlen(gs) < 30 {
		h.health = h.health1
	} else {
		h.health = h.health2
	}
	return h
}

func (h *geneHub) init() {
	m1 := [26]*mGenes{}
	m2 := [26][26]*mGenes{}
	m3 := [26][26][26]*mGenes{}
	for _, g := range h.genes {
		b := []byte(g.dna)
		b0 := b[0]
		if len(g.dna) == 1 {
			m1[b0-'a'] = m1[b0-'a'].add(g)
		} else if len(g.dna) == 2 {
			b1 := b[1]
			m2[b0-'a'][b1-'a'] = m2[b0-'a'][b1-'a'].add(g)
		} else if len(g.dna) > 2 {
			b1 := b[1]
			b2 := b[2]
			m3[b0-'a'][b1-'a'][b2-'a'] = m3[b0-'a'][b1-'a'][b2-'a'].add(g)
		}
	}
	h.m1 = m1
	h.m2 = m2
	h.m3 = m3
}

func (h *geneHub) sub(i, j int) *geneHub {
	h.start = i
	h.end = j
	return h
}

func maxlen(gs []gene) int {
	max := 0
	for _, g := range gs {
		if len(g.dna) > max {
			max = len(g.dna)
		}
	}
	return max
}

func (h *geneHub) health1(d dna) int {
	n := len(d)
	r := 0
	for i := 0; i < n; i++ {
		bi := d[i]

		m1i := h.m1[bi-'a']
		if m1i != nil {
			for _, l := range m1i.lens() {
				if i < n {
					pg := m1i.get(d[i : i+l])
					if pg != nil {
						for _, p := range pg.pairs {
							if p.idx >= h.start && p.idx <= h.end {
								r += p.val
							}
						}
						break
					}
				}
			}
		}
		if i+1 < n {
			bi1 := d[i+1]
			m2i := h.m2[bi-'a'][bi1-'a']
			if m2i != nil {
				for _, l := range m2i.lens() {
					p := i + l
					if p <= n {
						pg := m2i.get(d[i:p])
						if pg != nil {
							for _, p := range pg.pairs {
								if p.idx >= h.start && p.idx <= h.end {
									r += p.val
								}
							}
							break
						}
					}
				}
			}
		}
		if i+2 < n {
			bi1 := d[i+1]
			bi2 := d[i+2]
			m3i := h.m3[bi-'a'][bi1-'a'][bi2-'a']
			if m3i != nil {
				for _, l := range m3i.lens() {
					p := i + l
					if p <= n {
						pg := m3i.get(d[i:p])
						if pg != nil {
							for _, p := range pg.pairs {
								if p.idx >= h.start && p.idx <= h.end {
									r += p.val
								}
							}
						}
					}
				}
			}
		}
	}
	return r
}

func (h *geneHub) health2(d dna) int {
	n := len(d)
	r := 0
	for i := 0; i < n; i++ {
		bi := d[i]

		m1i := h.m1[bi-'a']
		if m1i != nil {
			for _, pg := range m1i.sm {
				if i < n {
					if d[i:i+1] == pg.dna {
						for _, p := range pg.pairs {
							if p.idx >= h.start && p.idx <= h.end {
								r += p.val
							}
						}
						break
					}
				}
			}
		}
		if i+1 < n {
			bi1 := d[i+1]
			m2i := h.m2[bi-'a'][bi1-'a']
			if m2i != nil {
				for _, pg := range m2i.sm {
					p := i + len(pg.dna)
					if p <= n {
						if d[i:i+2] == pg.dna {
							for _, p := range pg.pairs {
								if p.idx >= h.start && p.idx <= h.end {
									r += p.val
								}
							}
							break
						}
					}
				}
			}
		}
		if i+2 < n {
			bi1 := d[i+1]
			bi2 := d[i+2]
			m3i := h.m3[bi-'a'][bi1-'a'][bi2-'a']
			if m3i != nil {
				for _, pg := range m3i.sm {
					p := i + len(pg.dna)
					if p <= n {
						if d[i:p] == pg.dna {
							for _, p := range pg.pairs {
								if p.idx >= h.start && p.idx <= h.end {
									r += p.val
								}
							}
						}
					}
				}
			}
		}
	}
	return r
}

type dna string
