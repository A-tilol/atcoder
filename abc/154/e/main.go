package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var dp = make(map[string]map[int]map[bool]int)
var n string
var k int

// algorithm: dp, 桁dp
func main() {
	n, k = readString(), readInt()

	fmt.Println(dfs(n, 0, true))

	// fmt.Println(dp)
}

func dfs(num string, cnt int, isFull bool) int {
	// fmt.Println(num, cnt, isFull)
	if cnt > k {
		return 0
	}

	if v, ok := dp[num][cnt][isFull]; ok {
		return v
	}

	if _, ok := dp[num]; !ok {
		dp[num] = make(map[int]map[bool]int)
	}
	if _, ok := dp[num][cnt]; !ok {
		dp[num][cnt] = make(map[bool]int)
	}

	if len(num) == 1 {
		if cnt == k {
			dp[num][cnt][isFull] = 1
		} else if cnt+1 == k {
			t, _ := strconv.Atoi(num)
			dp[num][cnt][isFull] = t
		} else {
			dp[num][cnt][isFull] = 0
		}
		return dp[num][cnt][isFull]
	}

	ans := 0
	next := num[1:]
	top, _ := strconv.Atoi(num[0:1])
	if isFull && top == 0 {
		dp[num][cnt][isFull] = dfs(next, cnt, true)
		return dp[num][cnt][isFull]
	}

	if isFull {
		ans += dfs(next, cnt+1, true)
		ans += dfs("9"+next[1:], cnt+1, false) * (top - 1) // except zero and the largest num
		ans += dfs("9"+next[1:], cnt, false)               // zero
	} else {
		ans += dfs("9"+next[1:], cnt+1, false) * 9
		ans += dfs("9"+next[1:], cnt, false)
	}

	dp[num][cnt][false] = ans
	return dp[num][cnt][false]
}

func comb(n, r int) int {
	if r == 0 {
		return 1
	}
	return comb(n, r-1) * (n - r + 1) / r
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
