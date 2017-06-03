package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	t := readInput(os.Stdin)
	fmt.Println(counter2(t))
}

func readInput(r io.Reader) (t int) {
	fmt.Fscanf(r, "%d\n", &t)
	return
}

var start = 3

type pair struct{ time, value int }

func (p *pair) String() string {
	return fmt.Sprintf("[%d, %d]", p.time, p.value)
}

func counter() <-chan *pair {
	c := make(chan *pair)
	go func() {
		curMark := start
		t := 1
		curT := t
		var v int
		for {
			v = curMark - curT + 1
			c <- &pair{t, v}
			t++
			if v == 1 {
				curMark *= 2
				curT = 1
			} else {
				curT++
			}
		}
	}()
	return c
}

func counter2(t int) int {
	curMark := start
	for t > curMark {
		t -= curMark
		curMark *= 2
	}
	return curMark - t + 1
}

func counter3() <-chan int {
	c := make(chan int)
	go func() {
		t := 1
		for {
			c <- counter2(t)
			t++
		}
	}()
	return c
}

func counter4() <-chan *pair {
	c := make(chan *pair)
	go func() {
		t := 1
		for {
			c <- &pair{t, counter2(t)}
			t++
		}
	}()
	return c
}

func counter5() <-chan int {
	c := make(chan int)
	go func() {
		curMark := start
		curT := 1
		var v int
		for {
			v = curMark - curT + 1
			c <- v
			if v == 1 {
				curMark *= 2
				curT = 1
			} else {
				curT++
			}
		}
	}()
	return c
}
