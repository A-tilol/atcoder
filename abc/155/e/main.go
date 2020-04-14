package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var n int
var num []int

// algorithm: dp, 桁dp
func main() {
	s := readBytes()
	n = len(s)

	num = make([]int, n)
	for i := range s {
		num[i] = int(s[i] - '0')
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][1] = 1

	for i := range num {
		dp[i+1][1] = min(
			dp[i][1]+(10-num[i]-1),
			dp[i][0]+num[i]+1)

		dp[i+1][0] = min(
			dp[i][1]+(10-num[i]),
			dp[i][0]+num[i])
	}

	fmt.Println(dp[n][0])
}

// byte -> string -> int の変換がちょっと遅い
func main3() {
	s := readBytes()
	n = len(s)

	num = make([]int, n)
	for i := range s {
		num[i], _ = strconv.Atoi(string(s[i]))
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][1] = 1

	for i := range num {
		dp[i+1][1] = min(
			dp[i][1]+(10-num[i]-1),
			dp[i][0]+num[i]+1)

		dp[i+1][0] = min(
			dp[i][1]+(10-num[i]),
			dp[i][0]+num[i])
	}

	fmt.Println(dp[n][0])
}

// 毎回mapの初期化処理が入る
func main2() {
	s := readBytes()
	n = len(s)

	num = make([]int, n)
	for i := range s {
		num[i], _ = strconv.Atoi(string(s[i]))
	}

	dp1[-1] = make(map[bool]int)
	dp1[-1][true] = 1
	for i := range num {
		dp1[i] = make(map[bool]int)

		dp1[i][true] = min(
			dp1[i-1][true]+(10-num[i]-1),
			dp1[i-1][false]+num[i]+1)

		dp1[i][false] = min(
			dp1[i-1][true]+(10-num[i]),
			dp1[i-1][false]+num[i])
	}

	fmt.Println(dp1[n-1][false])
}

// TLE
func main1() {
	s := readBytes()
	n = len(s)

	num = make([]int, n)
	for i := range s {
		num[i], _ = strconv.Atoi(string(s[i]))
	}

	fmt.Println(min(dfs(0, true)+1, dfs(0, false)))
}

var dp1 = make(map[int]map[bool]int)

func dfs(i int, ten bool) int {
	if v, ok := dp1[i][ten]; ok {
		return v
	}

	if _, ok := dp1[i]; !ok {
		dp1[i] = make(map[bool]int)
	}

	if i == n-1 {
		if ten {
			dp1[i][ten] = 10 - num[i]
		} else {
			dp1[i][ten] = num[i]
		}
		// fmt.Println("end:", i, ten, dp1[i][ten])
		return dp1[i][ten]
	}

	ans := 0
	if ten {
		ans = dfs(i+1, true) + (10 - num[i] - 1)
		// fmt.Println(1, i, ten, ans)
		ans = min(ans, dfs(i+1, false)+(10-num[i]))
		// fmt.Println(2, i, ten, dfs(i+1, false)+(10-num[i]))
	} else {
		ans = dfs(i+1, true) + num[i] + 1
		// fmt.Println(3, i, ten, ans)
		ans = min(ans, dfs(i+1, false)+num[i])
		// fmt.Println(4, i, ten, dfs(i+1, false)+num[i])
	}

	dp1[i][ten] = ans
	return dp1[i][ten]
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
