package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

// ソート済配列arrに含まれる数のうちvを超えない最大の数の配列のインデックスを返す
func bisectLeft(arr []int, v, l, r int) int {
	if v < arr[l] {
		return 0
	}

	var m int
	for l < r {
		m = (l+r)/2 + 1
		if arr[m] <= v {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}

func main() {
	n, m := readInt(), readInt()
	h := make([]int, n)
	for i := range h {
		h[i] = readInt()
	}
	w := make([]int, m)
	for i := range w {
		w[i] = readInt()
	}

	sort.Ints(h)

	cum1 := make([]int, n/2)
	cum2 := make([]int, n/2)
	for i := range cum1 {
		cum1[i] = h[2*i+1] - h[2*i]
		cum2[i] = h[2*i+2] - h[2*i+1]
		if i > 0 {
			cum1[i] += cum1[i-1]
			cum2[i] += cum2[i-1]
		}
	}

	if n/2 == 0 {
		cum1 = []int{0}
		cum2 = []int{0}
	}

	ans := math.MaxInt64
	for i := range w {
		idx := bisectLeft(h, w[i], 0, len(h)-1)

		tmp := 0
		if idx != 0 {
			tmp = cum1[min(len(cum1)-1, (idx-1)/2)]
		}

		if idx%2 == 1 {
			tmp += abs(h[min(len(h)-1, idx+1)] - w[i])
			if idx/2 < len(cum2) {
				tmp += cum2[len(cum2)-1] - cum2[idx/2]
			}
		} else {
			tmp += abs(w[i] - h[idx])
			tmp += cum2[len(cum2)-1]
			if idx != 0 && idx/2-1 < len(cum2) {
				tmp -= cum2[idx/2-1]
			}
		}

		ans = min(ans, tmp)
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
