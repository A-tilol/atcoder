package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

// algorithm: bit全探索
func main() {
	n := readInt()

	m := make([][]byte, n)
	for i := range m {
		m[i] = make([]byte, n)
	}

	for i := 0; i < n; i++ {
		a := readInt()

		for j := 0; j < a; j++ {
			x, y := readInt()-1, readInt()
			if y == 1 {
				m[i][x] = '1'
			} else {
				m[i][x] = '0'
			}
		}
	}

	all := int(math.Pow(2, float64(n)))
	format := "%0" + strconv.Itoa(n) + "b"
	ans := 0
	for i := 0; i < all; i++ {
		b := []byte(fmt.Sprintf(format, i))
		honest := 0

	LABEL:
		for k := range m {
			if b[k] == '0' {
				continue
			}
			honest++

			for l := range m[k] {
				if m[k][l] == 0 {
					continue
				}
				if m[k][l] != b[l] {
					honest = 0
					break LABEL
				}

			}
		}

		ans = max(ans, honest)
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
