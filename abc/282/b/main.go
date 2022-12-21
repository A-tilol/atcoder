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
	s := make([][]byte, n)
	for i := range s {
		s[i] = readBytes()
	}

	ans := 0
	for i := range s {
		x_inds := []int{}
		for j, c := range s[i] {
			if c == 'x' {
				x_inds = append(x_inds, j)
			}
		}
		if len(x_inds) == 0 {
			ans += m - (i + 1)
			continue
		}

		for j := i + 1; j < n; j++ {
			ok := true
			for _, ind := range x_inds {
				if s[j][ind] == 'x' {
					ok = false
					break
				}
			}
			if ok {
				ans++
			}
		}
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
