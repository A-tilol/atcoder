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
	h, w := readInt(), readInt()

	whiteCnt := 0
	s := make([][]byte, h)
	visited := make([][]bool, h)
	for i := range s {
		s[i] = readBytes()
		visited[i] = make([]bool, w)
		for j := range s[i] {
			if s[i][j] == '.' {
				whiteCnt++
			}
		}
	}

	q := []point{point{0, 0, 0}}
	visited[0][0] = true
	nexts := []point{
		point{1, 0, 0},
		point{-1, 0, 0},
		point{0, 1, 0},
		point{0, -1, 0},
	}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for i := range nexts {
			ny := p.y + nexts[i].y
			nx := p.x + nexts[i].x
			ncnt := p.cnt + 1
			if ny < 0 || ny >= h || nx < 0 || nx >= w || s[ny][nx] == '#' {
				continue
			}
			if visited[ny][nx] {
				continue
			}
			if ny == h-1 && nx == w-1 {
				fmt.Println(whiteCnt - ncnt - 1)
				return
			}
			visited[ny][nx] = true
			q = append(q, point{ny, nx, ncnt})
		}
	}

	fmt.Println(-1)
}

type point struct {
	y   int
	x   int
	cnt int
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
