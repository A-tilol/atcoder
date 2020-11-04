package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var toPrime []int

func findPrimes(n int) {
	for i := range toPrime {
		toPrime[i] = 1
	}
	for i := n; i >= 2; i-- {
		for j := i + i; j <= n; j += i {
			toPrime[j] = i
		}
	}
}

func main() {
	n := readInt()

	toPrime = make([]int, n+1)
	findPrimes(n)

	ans := 1
	for i := 2; i < n; i++ {
		primeTocnt := map[int]int{}
		num := i
		for {
			if toPrime[num] == 1 {
				primeTocnt[num]++
				break
			}
			primeTocnt[toPrime[num]]++
			num /= toPrime[num]
		}

		tmp := 1
		for _, v := range primeTocnt {
			tmp *= v + 1
		}
		// fmt.Println(i, primeTocnt, tmp)
		ans += tmp
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
