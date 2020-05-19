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
	s := readBytes()
	q := readInt()

	b := make([]byte, len(s)+2*q+10)
	start := q + 5
	end := start + len(s) - 1
	for i := range s {
		b[start+i] = s[i]
	}

	flipCnt := 0
	for i := 0; i < q; i++ {
		if readInt() == 1 {
			flipCnt++
			continue
		}

		f, c := readInt(), readBytes()[0]
		if (f == 1 && flipCnt%2 == 0) || (f == 2 && flipCnt%2 == 1) {
			start--
			b[start] = c
		} else {
			end++
			b[end] = c
		}
	}

	if flipCnt%2 == 1 {
		i, j := start, end
		for i < j {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		}
	}
	fmt.Println(string(b[start : end+1]))
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
