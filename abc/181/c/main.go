package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var e = 1e-5

func main() {
	n := readInt()

	x := make([]float64, n)
	y := make([]float64, n)
	for i := range x {
		x[i], y[i] = readFloat64(), readFloat64()
	}

	for i := range x {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if x[i] == x[j] {
					if x[i] == x[k] {
						fmt.Println("Yes")
						return
					}
					continue
				}
				if y[i] == y[j] {
					if y[i] == y[k] {
						fmt.Println("Yes")
						return
					}
					continue
				}
				a := (y[i] - y[j]) / (x[i] - x[j])
				b := y[i] - a*x[i]
				yk := a*x[k] + b
				if yk > y[k]-e && yk < y[k]+e {
					fmt.Println("Yes")
					return
				}
			}
		}
	}

	fmt.Println("No")
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
