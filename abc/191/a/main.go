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

func main() {
	debug.logging = true

	v, t, s, d := readInt(), readInt(), readInt(), readInt()

	if d >= t*v && d <= s*v {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
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
