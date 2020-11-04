package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	n, x, m := readInt(), readInt(), readInt()

	amap := map[int]int{}
	a := make([]int, 0)

	cur := x
	ans := 0
	for i := 0; i < n; i++ {
		// fmt.Println(i, cur, ans)
		if cur == 0 {
			//sum(0, i) ai
			break
		}
		if start, ok := amap[cur]; ok {
			// sum(0, i-1) ai * n / i + r
			cycleSum := 0
			for _, ai := range a[start:] {
				cycleSum += ai
			}
			// fmt.Println("cycleSum", cycleSum)

			cycleLen := (i - 1) - start + 1
			ans += cycleSum * ((n-start)/cycleLen - 1)

			// fmt.Println("ans excluded reminder", ans)

			reminder := (n - start) % cycleLen
			for j := 0; j < reminder; j++ {
				ans += a[start+j]
			}

			break
		}
		ans += cur
		amap[cur] = i
		a = append(a, cur)

		cur = cur * cur % m
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
