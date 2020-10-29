package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var h, w, ch, cw, dh, dw int
var s [][]byte
var area [][]int
var areaToDist = map[int]int{}

type Next struct {
	h int
	w int
}

var nexts []Next

func markAreaNum(i, j, num int) {
	if i < 0 || i >= h || j < 0 || j >= w {
		return
	}
	if area[i][j] != -1 || s[i][j] == '#' {
		return
	}
	area[i][j] = num

	markAreaNum(i+1, j, num)
	markAreaNum(i-1, j, num)
	markAreaNum(i, j+1, num)
	markAreaNum(i, j-1, num)
}

func findNextArea(i, j, areaNum int) {
	if i < 0 || i >= h || j < 0 || j >= w {
		return
	}
	if area[i][j] != areaNum {
		return
	}
	dist := areaToDist[area[i][j]]
	area[i][j] = -1 // visited

	for ii := i - 2; ii <= i+2; ii++ {
		for jj := j - 2; jj <= j+2; jj++ {
			if ii < 0 || ii >= h || jj < 0 || jj >= w {
				continue
			}
			if area[ii][jj] == areaNum || area[ii][jj] == -1 {
				continue
			}
			if _, ok := areaToDist[area[ii][jj]]; ok {
				continue
			}
			areaToDist[area[ii][jj]] = dist + 1
			nexts = append(nexts, Next{ii, jj})
		}
	}

	findNextArea(i+1, j, areaNum)
	findNextArea(i-1, j, areaNum)
	findNextArea(i, j+1, areaNum)
	findNextArea(i, j-1, areaNum)
}

func main() {
	h, w = readInt(), readInt()
	ch, cw = readInt()-1, readInt()-1
	dh, dw = readInt()-1, readInt()-1
	s = make([][]byte, h)
	for i := range s {
		s[i] = readBytes()
	}

	area = make([][]int, h)
	for i := range area {
		area[i] = make([]int, w)
		for j := range area[i] {
			area[i][j] = -1
		}
	}
	areaNum := 1
	for i := range area {
		for j := range area[i] {
			if area[i][j] != -1 || s[i][j] == '#' {
				continue
			}
			markAreaNum(i, j, areaNum)
			areaNum++
		}
	}

	nexts = make([]Next, 1)
	nexts[0] = Next{ch, cw}
	areaToDist[area[ch][cw]] = 0

	goalArea := area[dh][dw]
	for len(nexts) > 0 {
		next := nexts[0]
		nexts = nexts[1:]
		if area[next.h][next.w] == -1 {
			continue
		}
		findNextArea(next.h, next.w, area[next.h][next.w])
	}

	if dist, ok := areaToDist[goalArea]; ok {
		fmt.Println(dist)
	} else {
		fmt.Println(-1)
	}
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
