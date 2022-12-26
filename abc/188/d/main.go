package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

type pair struct {
	day  int
	cost int
}

func main() {
	n := readInt()
	c := readInt()

	arr := make([]pair, n*2)
	for i := 0; i < n; i++ {
		a, b, c := readInt(), readInt(), readInt()
		arr[2*i] = pair{a, c}
		arr[2*i+1] = pair{b + 1, -c}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].day < arr[j].day
	})

	ans, cur, i, day := 0, 0, 0, arr[0].day
	for day != arr[len(arr)-1].day {
		for day == arr[i].day {
			cur += arr[i].cost
			i++
		}
		if c < cur {
			ans += c * (arr[i].day - day)
		} else {
			ans += cur * (arr[i].day - day)
		}
		day = arr[i].day
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
