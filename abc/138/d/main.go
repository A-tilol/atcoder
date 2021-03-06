package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var n int
var ans, nodex []int
var amat [][]int

func main() {
	n = readInt()
	q := readInt()

	ans = make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	amat = make([][]int, n)
	for i := range amat {
		amat[i] = make([]int, 0)
	}

	for i := 0; i < n-1; i++ {
		a, b := readInt()-1, readInt()-1
		amat[a] = append(amat[a], b)
		amat[b] = append(amat[b], a)
	}

	nodex = make([]int, n)
	for i := 0; i < q; i++ {
		p, x := readInt()-1, readInt()
		nodex[p] += x
	}

	dfs(0, 0)

	for i := range ans {
		fmt.Println(ans[i])
	}
}

func dfs(i, x int) {
	if i == n {
		return
	}
	x += nodex[i]
	ans[i] = x
	for j := range amat[i] {
		if ans[amat[i][j]] == -1 {
			dfs(amat[i][j], x)
		}
	}
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
