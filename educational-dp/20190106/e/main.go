package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

func minInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func maxInt64(a, b int64) int64 {
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

func absInt64(a int64) int64 {
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

func sumInt64(a []int64) int64 {
	var ret int64
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

func gcdInt64(m, n int64) int64 {
	for m%n != 0 {
		m, n = n, m%n
	}
	return n
}

func lcmInt64(m, n int64) int64 {
	return m / gcdInt64(m, n) * n
}

// sort ------------------------------------------------------------

type int64Array []int64

func (s int64Array) Len() int           { return len(s) }
func (s int64Array) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s int64Array) Less(i, j int) bool { return s[i] < s[j] }

type xxx struct {
	x int
}

type sortArray []xxx

func (s sortArray) Len() int           { return len(s) }
func (s sortArray) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortArray) Less(i, j int) bool { return s[i].x < s[j].x }

// -----------------------------------------------------------------
var n, vmax int
var w int64
var items []item
var dp [][]int64

type item struct {
	w int64
	v int
}

func main() {
	n = readInt()
	w = readInt64()
	vmax = n * 1e3
	items = make([]item, n)
	for i := range items {
		items[i] = item{
			w: readInt64(),
			v: readInt(),
		}
	}
	dp = make([][]int64, n+1)
	for i := range dp {
		dp[i] = make([]int64, vmax+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt64
		}
	}
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 0; j <= vmax; j++ {
			if j-items[i-1].v >= 0 && dp[i-1][j-items[i-1].v] != math.MaxInt64 {
				dp[i][j] = minInt64(dp[i-1][j], dp[i-1][j-items[i-1].v]+items[i-1].w)
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	var ans int
	for i := range dp[0] {
		if dp[n][i] <= w {
			ans = i
		}
	}

	fmt.Println(ans)
}