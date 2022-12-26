package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type pair struct {
	a int
	b int
}

func main() {
	n := readInt()
	l := 1 << (n - 1)
	num, v := -1, -1
	for i := 0; i < l; i++ {
		a := readInt()
		if a > v {
			num, v = i+1, a
		}
	}
	num2, v2 := -1, -1
	for i := 0; i < l; i++ {
		a := readInt()
		if a > v2 {
			num2, v2 = l+i+1, a
		}
	}

	if v < v2 {
		fmt.Println(num)
	} else {
		fmt.Println(num2)
	}
}

func main2() {
	n := readInt()
	m := 1 << n
	a := make([]int, m)
	for i := 0; i < m; i++ {
		a[i] = readInt()
	}

	for k := 0; k < n-1; k++ {
		ind := -1
		for i := range a {
			if a[i] == -1 {
				continue
			}
			if ind == -1 {
				ind = i
			} else {
				if a[ind] > a[i] {
					a[i] = -1
				} else {
					a[ind] = -1
				}
				ind = -1
			}
		}
	}

	ind := -1
	for i := range a {
		if a[i] == -1 {
			continue
		}
		if ind == -1 {
			ind = i
		} else {
			if a[ind] < a[i] {
				fmt.Println(ind + 1)
			} else {
				fmt.Println(i + 1)
			}
			return
		}
	}
}

func main1() {
	n := readInt()
	m := 1 << n
	a := make([][]int, m)
	aa := make([][]int, m)
	for i := 0; i < m; i++ {
		a[i] = []int{i + 1, readInt()}
	}
	copy(aa, a)

	for len(aa) > 2 {
		a = aa
		aa = make([][]int, len(a)/2)
		for i := 0; i < len(a)/2; i++ {
			if a[2*i][1] > a[2*i+1][1] {
				aa[i] = a[2*i]
			} else {
				aa[i] = a[2*i+1]
			}
		}
	}

	if aa[0][1] > aa[1][1] {
		fmt.Println(aa[1][0])
	} else {
		fmt.Println(aa[0][0])
	}
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
