package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var mat = make(map[int][]int)

func main() {
	n, m := readInt(), readInt()

	for i := 0; i < m; i++ {
		a, b := readInt(), readInt()
		if _, ok := mat[a]; !ok {
			mat[a] = make([]int, 0)
		}
		mat[a] = append(mat[a], b)
		if _, ok := mat[b]; !ok {
			mat[b] = make([]int, 0)
		}
		mat[b] = append(mat[b], a)
	}

	nexts := make([]move, len(mat[1]))
	for i := range mat[1] {
		nexts[i] = move{mat[1][i], 1}
	}

	visited := map[int]bool{1: true}
	ans := make([]int, n)
	for len(nexts) > 0 && len(visited) < n {
		tmpNexts := make([]move, 0)
		// fmt.Println(nexts)
		for i := range nexts {
			if visited[nexts[i].cur] {
				continue
			}
			visited[nexts[i].cur] = true

			ans[nexts[i].cur-1] = nexts[i].pre

			if _, ok := mat[nexts[i].cur]; !ok {
				continue
			}
			for j := range mat[nexts[i].cur] {
				tmpNexts = append(tmpNexts, move{mat[nexts[i].cur][j], nexts[i].cur})
			}
		}
		nexts = tmpNexts
	}

	fmt.Println("Yes")
	for i := range ans[1:] {
		fmt.Println(ans[i+1])
	}
}

type move struct {
	cur int
	pre int
}

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
