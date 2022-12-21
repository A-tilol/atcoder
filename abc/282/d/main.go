package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	n, m := readInt(), readInt()
	ad := make([][]int, n)
	for i := range ad {
		ad[i] = []int{}
	}

	for i := 0; i < m; i++ {
		u, v := readInt()-1, readInt()-1
		ad[u] = append(ad[u], v)
		ad[v] = append(ad[v], u)
	}

	ans := n*(n-1)/2 - m
	colors := make([]bool, n)
	visited := make([]bool, n)
	for i := range ad {
		if visited[i] {
			continue
		}

		black, white := 0, 0
		que := []int{i}
		for len(que) > 0 {
			parent := que[0]
			que = que[1:]
			for _, child := range ad[parent] {
				if visited[child] {
					// Detect a odd number loop. Not bipartite graph.
					if colors[child] == colors[parent] {
						fmt.Println(0)
						return
					}
					continue
				}
				colors[child] = !colors[parent]
				visited[child] = true
				que = append(que, child)
				if colors[child] {
					black++
				} else {
					white++
				}
			}
		}
		ans -= black * (black - 1) / 2
		ans -= white * (white - 1) / 2
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
