package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var mat [][]int
var visited []bool
var ans []int
var outindex map[string]int
var k int

func dfs(cur, color int) {
	visited[cur] = true

	nexts := mat[cur]
	i, nxcolor := 0, 1
	for i < len(nexts) {
		if visited[nexts[i]] {
			i++
			continue
		}
		if nxcolor == color {
			nxcolor++
			continue
		}

		dfs(nexts[i], nxcolor)

		k = max(k, nxcolor)
		outi := outindex[strconv.Itoa(cur)+"-"+strconv.Itoa(nexts[i])]
		ans[outi] = nxcolor

		i++
		nxcolor++
	}
}

func main() {
	n := readInt()

	mat = make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, 0)
	}

	outindex = make(map[string]int)
	for i := 0; i < n-1; i++ {
		a, b := readInt()-1, readInt()-1
		mat[a] = append(mat[a], b)
		mat[b] = append(mat[b], a)

		aa, bb := strconv.Itoa(a), strconv.Itoa(b)
		outindex[aa+"-"+bb] = i
		outindex[bb+"-"+aa] = i

	}

	ans = make([]int, n-1)
	visited = make([]bool, n)

	dfs(0, -1)

	fmt.Println(k)
	for i := range ans {
		fmt.Println(ans[i])
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
	log.SetFlags(log.Lshortfile)
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
