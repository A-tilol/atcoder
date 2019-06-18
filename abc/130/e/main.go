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

func main() {
	var mod int64 = 1000000000 + 7
	n := readInt64()
	m := readInt64()
	S := make([]int64, n)
	T := make([]int64, m)
	for i := range S {
		S[i] = readInt64()
	}
	for i := range T {
		T[i] = readInt64()
	}

	sum := make([][]int64, n+1)
	for i := range sum {
		sum[i] = make([]int64, m+1)
	}

	for i := range S {
		for j := range T {
			if S[i] == T[j] {
				sum[i+1][j+1] = (sum[i][j+1] + sum[i+1][j] + 1) % mod
			} else {
				sum[i+1][j+1] = (sum[i][j+1] + sum[i+1][j] - sum[i][j]) % mod
			}
		}
	}
	fmt.Println(sum[n][m] + 1)
}
