package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	readString func() string
	readBytes  func() []byte
)

func init() {
	readString, readBytes = newReadString(os.Stdin)
}

func newReadString(ior io.Reader) (func() string, func() []byte) {
	r := bufio.NewScanner(ior)
	r.Buffer(make([]byte, 1024), int(1e+11))
	r.Split(bufio.ScanWords)

	f1 := func() string {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Text()
	}
	f2 := func() []byte {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Bytes()
	}
	return f1, f2
}

func readInt() int {
	return int(readInt64())
}

func readInt64() int64 {
	i, err := strconv.ParseInt(readString(), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	return i
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
	if a > 0 {
		return a
	}
	return -a
}

// sort ------------------------------------------------------------

type xxx struct {
	x int
}

type sortArray []xxx

func (s sortArray) Len() int           { return len(s) }
func (s sortArray) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortArray) Less(i, j int) bool { return s[i].x < s[j].x }

// -----------------------------------------------------------------

func df(i, n int, on []int) {
	if i == n {
		for i := range S {
			cnt := 0
			for j := range S[i] {
				if on[S[i][j]-1] == 1 {
					cnt++
				}
			}
			if cnt%2 != P[i] {
				return
			}
		}
		ans++
		return
	}
	df(i+1, n, append(on, 0))
	df(i+1, n, append(on, 1))
}

var ans int
var S [][]int
var P []int

func main() {
	n := readInt()
	m := readInt()

	S = make([][]int, m)
	for i := range S {
		k := readInt()
		S[i] = make([]int, k)
		for j := range S[i] {
			S[i][j] = readInt()
		}
	}

	P = make([]int, m)
	for i := range P {
		P[i] = readInt()
	}

	df(0, n, make([]int, 0, n))

	fmt.Println(ans)
}
