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

func main() {
	n, d, a := readInt(), readInt(), readInt()
	m := make([]xxx, n)
	for i := range m {
		m[i] = xxx{readInt(), readInt()}
	}

	sort.Sort(sortArray(m))

	itoj := make(map[int]int)
	j := 0
	for i := range m {
		for j+1 < n && m[j+1].x <= m[i].x+2*d {
			j++
		}
		itoj[i] = j
	}

	decmap := make(map[int]int)
	ans, sum := 0, 0
	for i := range m {
		if v, ok := decmap[i]; ok {
			sum -= v
		}
		if m[i].h-sum <= 0 {
			continue
		}
		cnt := (m[i].h - sum + a - 1) / a
		ans += cnt
		sum += cnt * a
		decmap[itoj[i]+1] += cnt * a
	}

	fmt.Println(ans)
}

/*
TLE after_contest_01
*/
func main1() {
	n, d, a := readInt(), readInt(), readInt()
	m := make([]xxx, n)
	for i := range m {
		m[i] = xxx{readInt(), readInt()}
	}

	sort.Sort(sortArray(m))

	tar := make([]xxx, 0)
	ans := 0
	for i := 0; i < n; i++ {
		tar = append(tar, m[i])

		if i == n-1 {
			maxh := 0
			for j := range tar {
				maxh = max(maxh, tar[j].h)
			}
			ans += (maxh + a - 1) / a
			break
		}

		for m[i+1].x > tar[0].x+2*d {
			cnt := (tar[0].h + a - 1) / a
			ans += cnt
			dech := cnt * a
			newtar := make([]xxx, 0)
			for j := range tar {
				if tar[j].h > dech {
					newtar = append(newtar, xxx{tar[j].x, tar[j].h - dech})
				}
			}
			tar = newtar
			if len(tar) == 0 {
				break
			}
		}
	}

	fmt.Println(ans)
}

// sort ------------------------------------------------------------

type xxx struct {
	x int
	h int
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
