package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// 繰り返し2乗法,  a^n mod を計算する
func modpow(a, n int) int {
	res := 1
	for n > 0 {
		if n&1 != 0 {
			res = res * a % mod
		}
		a = a * a % mod
		n >>= 1
	}
	return res
}

var mod = int(1e9) + 7
var n int
var k = int(1e5*2 + 1)
var inv = make([]int, k)
var comb = make([]int, k)

// combination
//   condition: about n >= 1e8, k <= 1e7
func initComb() {
	inv[1] = 1
	for i := 2; i < k; i++ {
		inv[i] = mod - inv[mod%i]*(mod/i)%mod
	}

	comb[0] = 1
	for i := 1; i < k; i++ {
		comb[i] = comb[i-1] * ((n - i + 1) * inv[i] % mod) % mod
	}
}

// algorithm: 繰り返し2乗法, combination
func main() {
	n = readInt()
	a, b := readInt(), readInt()

	a = min(a, n-a)
	b = min(b, n-b)

	initComb()

	ans := modpow(2, n) - 1 - comb[a] - comb[b] + 3*mod
	fmt.Println(ans % mod)
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
