package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"runtime"
	"strconv"
)

func baseNTo10(x []int, n int) int {
	baseN := 0
	r := 1
	for i := len(x) - 1; i >= 0; i-- {
		rem := x[i] % 10
		if rem != 0 && math.MaxInt64/rem < r {
			return math.MaxInt64
		}
		if math.MaxInt64-baseN < rem*r {
			return math.MaxInt64
		}
		baseN += x[i] % 10 * r
		if i == 0 {
			break
		}
		if math.MaxInt64/n < r {
			return math.MaxInt64
		}
		r *= n
	}
	return baseN
}

func main() {
	debug.logging = false

	s, m := readString(), readInt()
	if len(s) == 1 {
		xx, _ := strconv.Atoi(s)
		if xx <= m {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
		return
	}

	d := 0
	x := make([]int, len(s))
	for i := range s {
		digit, _ := strconv.Atoi(s[i : i+1])
		d = max(d, digit)
		x[i] = digit
	}

	l, r := d, m
	for l < r {
		mid := (l+r)/2 + 1
		base10 := baseNTo10(x, mid)
		if base10 > m {
			r = mid - 1
		} else {
			l = mid
		}
	}
	debug.println(l, r)

	fmt.Println(max(0, r-d))
}

var readString func() string
var debug *debugLogger

func init() {
	log.SetFlags(log.Lshortfile)
	readString = newReadString(os.Stdin)
	debug = newDebugLogger(true)
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

func readInts(n int) []int {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = readInt()
	}
	return arr
}

func readFloat64() float64 {
	f, err := strconv.ParseFloat(readString(), 64)
	if err != nil {
		panic(err.Error())
	}
	return f
}

func readFloat64s(n int) []float64 {
	arr := make([]float64, n)
	for i := range arr {
		arr[i] = readFloat64()
	}
	return arr
}

func min(a ...int) int {
	x := math.MaxInt64
	for i := range a {
		if x > a[i] {
			x = a[i]
		}
	}
	return x
}

func max(a ...int) int {
	x := math.MinInt64
	for i := range a {
		if x < a[i] {
			x = a[i]
		}
	}
	return x
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

type debugLogger struct {
	logging bool
}

func newDebugLogger(logging bool) *debugLogger {
	return &debugLogger{logging: logging}
}

func (d *debugLogger) print(args ...interface{}) {
	if d.logging {
		fmt.Print(args...)
	}
}

func (d *debugLogger) printf(format string, args ...interface{}) {
	if d.logging {
		_, _, line, _ := runtime.Caller(1)
		format = fmt.Sprintf("(l%d) ", line) + format
		fmt.Printf(format, args...)
	}
}

func (d *debugLogger) println(args ...interface{}) {
	if d.logging {
		_, _, line, _ := runtime.Caller(1)
		format := fmt.Sprintf("(l%d)", line)
		args = append([]interface{}{format}, args...)
		fmt.Println(args...)
	}
}
