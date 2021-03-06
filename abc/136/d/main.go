package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

func main() {
	s := readString()

	rindices := make([]int, 0, len(s))
	lindices := make([]int, 0, len(s))
	for i := range s {
		if s[i] == 'R' {
			rindices = append(rindices, i)
		} else {
			lindices = append(lindices, i)
		}
	}

	ans := make([]int, len(s))
	for i := range s {
		if s[i] == 'R' {
			lidx := lindices[sort.SearchInts(lindices, i)]
			if (lidx-i)%2 == 0 {
				ans[lidx]++
			} else {
				ans[lidx-1]++
			}
		} else {
			ridx := rindices[sort.SearchInts(rindices, i)-1]
			if (i-ridx)%2 == 0 {
				ans[ridx]++
			} else {
				ans[ridx+1]++
			}
		}
	}

	for i := range ans {
		fmt.Printf("%v ", ans[i])
	}
}

func main1() {
	s := readString()

	ans := make([]int, len(s))
	prec := byte(' ')
	prej := 0
	for i := range s {
		c := s[i]
		j := 1

		if c == 'R' {
			if c == prec {
				j = prej - 1
			}
			for s[i+j] == 'R' {
				j++
			}
			if j%2 == 0 {
				ans[i+j]++
			} else {
				ans[i+j-1]++
			}
		} else {
			if c == prec {
				j = prej + 1
			}
			for s[i-j] == 'L' {
				j++
			}
			if j%2 == 0 {
				ans[i-j]++
			} else {
				ans[i-j+1]++
			}
		}
		prec, prej = c, j
	}

	for i := range ans {
		fmt.Printf("%v ", ans[i])
	}
}
