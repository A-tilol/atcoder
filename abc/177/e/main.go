package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const N = int(1e6)

var primes = make([]int, N+1)

func findPrimes() {
	for i := range primes {
		primes[i] = 1
	}

	for i := N; i > 1; i-- {
		for j := i; j <= N; j += i {
			primes[j] = i
		}
	}
}

func main() {
	findPrimes()

	n := readInt()
	a := make([]int, n)
	for i := range a {
		a[i] = readInt()
	}

	cnt := make([]int, N+1)
	for i := range a {
		ps := map[int]bool{}
		cur := a[i]
		p := primes[cur]
		for p != 1 {
			if _, ok := ps[p]; !ok {
				cnt[p]++
				ps[p] = true
			}
			cur /= p
			p = primes[cur]
		}
	}

	pairwise, setwise := true, true
	for i := range cnt {
		if cnt[i] > 1 {
			pairwise = false
		}
		if cnt[i] == n {
			setwise = false
		}
	}

	if pairwise {
		fmt.Println("pairwise coprime")
	} else if setwise {
		fmt.Println("setwise coprime")
	} else {
		fmt.Println("not coprime")
	}
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
