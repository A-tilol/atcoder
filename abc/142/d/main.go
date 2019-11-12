package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	a, b := readInt(), readInt()
	g := gcd(a, b)
	divisors := listDivisors(g)

	sort.Ints(divisors)

	ok := make([]bool, len(divisors))
	for i := range ok {
		ok[i] = true
	}

	for i := 1; i < len(divisors); i++ {
		end := true
		for j := i + 1; j < len(divisors); j++ {
			if !ok[j] {
				continue
			}
			if divisors[j]%divisors[i] == 0 {
				ok[j] = false
			}
			end = false
		}
		if end {
			break
		}
	}

	ans := 0
	for i := range ok {
		if ok[i] {
			ans++
		}
	}

	fmt.Println(ans)
}

func listDivisors(n int) []int {
	ret := make([]int, 0)
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i != 0 {
			continue
		}
		ret = append(ret, i)
		if n/i != i {
			ret = append(ret, n/i)
		}
	}
	return ret
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
