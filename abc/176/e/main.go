package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

type A struct {
	i   int
	tot int
}

type Pair struct {
	h int
	w int
}

func main() {
	h, w, m := readInt(), readInt(), readInt()
	yoko := make([]A, h)
	for i := range yoko {
		yoko[i].i = i
	}
	tate := make([]A, w)
	for i := range tate {
		tate[i].i = i
	}

	tar := make([]Pair, 0)
	for i := 0; i < m; i++ {
		hi, wi := readInt()-1, readInt()-1
		yoko[hi].tot++
		tate[wi].tot++

		tar = append(tar, Pair{hi, wi})
	}

	sort.Slice(yoko, func(i, j int) bool { return yoko[i].tot > yoko[j].tot })
	sort.Slice(tate, func(i, j int) bool { return tate[i].tot > tate[j].tot })

	hindicies := map[int]bool{}
	maxv := yoko[0].tot
	for i := range yoko {
		if yoko[i].tot != maxv {
			break
		}
		hindicies[yoko[i].i] = true
	}

	windicies := map[int]bool{}
	maxv = tate[0].tot
	for i := range tate {
		if tate[i].tot != maxv {
			break
		}
		windicies[tate[i].i] = true
	}

	cnt := 0
	for i := range tar {
		_, ok1 := hindicies[tar[i].h]
		_, ok2 := windicies[tar[i].w]
		if ok1 && ok2 {
			cnt++
		}
	}

	ans := len(hindicies) * len(windicies)
	if cnt == ans {
		ans--
	}
	fmt.Println(ans)
}

var readString func() string

func init() {
	log.SetFlags(log.Lshortfile)
	readString = newReadString(os.Stdin)
}

func newReadString(ior io.Reader) func() string {
	r := bufio.NewScanner(ior)
	r.Buffer(make([]byte, 1024), int(1e+11))
	r.Split(bufio.ScanWords)

	return func() string {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Text()
	}
}

func readBytes() []byte {
	return []byte(readString())
}

func readInt64() int64 {
	i, err := strconv.ParseInt(readString(), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	return i
}

func readInt() int {
	return int(readInt64())
}

func readFloat64() float64 {
	f, err := strconv.ParseFloat(readString(), 64)
	if err != nil {
		panic(err.Error())
	}
	return f
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sum(a []int) int {
	var ret int
	for i := range a {
		ret += a[i]
	}
	return ret
}

func sumFloat64(a []float64) float64 {
	var ret float64
	for i := range a {
		ret += a[i]
	}
	return ret
}

func gcd(m, n int) int {
	for m%n != 0 {
		m, n = n, m%n
	}
	return n
}

func lcm(m, n int) int {
	return m / gcd(m, n) * n
}
