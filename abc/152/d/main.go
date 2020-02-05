package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	n := readInt()
	a := make([][]int, 10)
	for i := range a {
		a[i] = make([]int, 10)
	}

	for i := 1; i < n+1; i++ {
		nb := []byte(strconv.Itoa(i))
		s, _ := strconv.Atoi(string(nb[0]))
		e, _ := strconv.Atoi(string(nb[len(nb)-1]))
		a[s][e]++
	}

	ans := 0
	for i := range a {
		for j := range a {
			ans += a[i][j] * a[j][i]
		}
	}
	fmt.Println(ans)
}

func main1() {
	n := readInt()
	nstr := strconv.Itoa(n)
	nlen := len(nstr)
	nstart, nmid, nend := n, 1, n
	if nlen > 1 {
		nstart, _ = strconv.Atoi(nstr[0:1])
		nend, _ = strconv.Atoi(nstr[nlen-1:])
	}
	if nlen > 2 {
		nmid, _ = strconv.Atoi(nstr[1 : nlen-1])
	}

	ans := 0
	for i := 1; i <= n; i++ {
		if i%10 == 0 {
			continue
		}

		a := strconv.Itoa(i)
		s, _ := strconv.Atoi(a[len(a)-1:])
		e, _ := strconv.Atoi(a[0:1])

		if s == e {
			ans++
		}

		for j := 0; j <= nlen-2; j++ {
			if j < nlen-2 {
				ans += int(math.Pow(10, float64(j)))
				continue
			}
			if s > nstart {
				break
			} else if s < nstart {
				ans += int(math.Pow(10, float64(j)))
			} else {
				if e <= nend {
					ans += nmid
				} else {
					ans += max(0, nmid-1)
				}
			}
		}
	}

	fmt.Println(ans)
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
