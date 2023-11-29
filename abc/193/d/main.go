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

	k, s, t := readInt(), readIntsOneByOne(), readIntsOneByOne()

	s_cnts := make([]int, 10)
	t_cnts := make([]int, 10)
	for i := range s[:4] {
		s_cnts[s[i]]++
		t_cnts[t[i]]++
	}

	s_score4, t_score4 := 0, 0
	for i := 1; i < 10; i++ {
		s_score4 += i * int(math.Pow10(s_cnts[i]))
		t_score4 += i * int(math.Pow10(t_cnts[i]))
	}
	debug.println(s_score4)
	debug.println(t_score4)

	s_scores := make([]int, 10)
	t_scores := make([]int, 10)
	for i := 1; i < 10; i++ {
		if k-s_cnts[i] > 0 {
			s_scores[i] = s_score4 - i*int(math.Pow10(s_cnts[i])) + i*int(math.Pow10(s_cnts[i]+1))
		}
		if k-t_cnts[i] > 0 {
			t_scores[i] = t_score4 - i*int(math.Pow10(t_cnts[i])) + i*int(math.Pow10(t_cnts[i]+1))
		}
	}
	debug.println(s_scores)
	debug.println(t_scores)

	ans := 0.0
	for i := 1; i < 10; i++ {
		if s_scores[i] == 0 {
			continue
		}
		for j := 1; j < 10; j++ {
			if t_scores[i] == 0 || s_scores[i] <= t_scores[j] {
				continue
			}
			ans += float64((k - s_cnts[i]) * (k - t_cnts[j]))
		}
	}
	ans /= float64((9*k - 4) * (9*k - 4))

	fmt.Println(ans)
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

func readIntsOneByOne() []int {
	s := readString()
	arr := make([]int, len(s))
	for i := range s {
		arr[i], _ = strconv.Atoi(s[i : i+1])
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
