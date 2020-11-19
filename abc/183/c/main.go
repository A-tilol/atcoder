package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var n, k, visitedAll int
var t [][]int

func dfs(total, from int, visited int) int {
	if visited == visitedAll {
		total += t[from][0]
		if total == k {
			return 1
		}
		return 0
	}

	ans := 0
	for i := 1; i < n; i++ {
		if ((1 << uint(i)) & visited) > 0 {
			continue
		}
		ans += dfs(total+t[from][i], i, visited+1<<uint(i))
	}

	return ans
}

func main() {
	n, k = readInt(), readInt()
	t = make([][]int, n)
	for i := range t {
		t[i] = make([]int, n)
		for j := range t[i] {
			t[i][j] = readInt()
		}
	}

	visited := 1
	for i := 0; i < n; i++ {
		visitedAll += 1 << uint(i)
	}

	fmt.Println(dfs(0, 0, visited))
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
