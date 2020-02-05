package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var h, w int
var s [][]byte
var wall byte = '#'
var visited map[int]struct{}

type pos struct {
	x int
	y int
	d int
}

var direction = []pos{
	pos{-1, 0, 0},
	pos{0, 1, 0},
	pos{1, 0, 0},
	pos{0, -1, 0},
}

func bfs(i, j int) int {
	move := 0
	q := []pos{pos{i, j, 0}}
	visited[i*w+j] = struct{}{}

	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		for _, d := range direction {
			nx := p.x + d.x
			ny := p.y + d.y
			nd := p.d + 1
			if _, ok := visited[nx*w+ny]; ok {
				continue
			}
			if nx < 0 || nx >= h || ny < 0 || ny >= w || s[nx][ny] == wall {
				continue
			}
			q = append(q, pos{nx, ny, nd})
			move = max(move, nd)
			visited[nx*w+ny] = struct{}{}
		}
	}

	return move
}

// algorithm: 幅優先探索(bfs)
func main() {
	h, w = readInt(), readInt()
	s = make([][]byte, h)
	for i := range s {
		s[i] = readBytes()
	}

	ans := 0
	for i := range s {
		for j := range s[i] {
			if s[i][j] == wall {
				continue
			}
			visited = make(map[int]struct{})
			ans = max(ans, bfs(i, j))
		}
	}

	fmt.Println(ans)
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
