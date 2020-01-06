package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

var cost = map[int]int{
	1: 2,
	2: 5,
	3: 5,
	4: 4,
	5: 5,
	6: 6,
	7: 3,
	8: 7,
	9: 6,
}

func main() {
	n, m := readInt(), readInt()

	matchNums := make([]matchNum, m)
	for i := range matchNums {
		a := readInt()
		matchNums[i] = matchNum{a, cost[a]}
	}

	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0

	for i := 0; i < n+1; i++ {
		for _, x := range matchNums {
			if x.cost+i > n {
				continue
			}
			dp[i+x.cost] = max(dp[i+x.cost], dp[i]+1)
		}
	}

	sort.Sort(sort.Reverse(sortArray(matchNums)))

	ans := ""
	j := n
	for i := dp[n]; i > 0; i-- {
		for _, x := range matchNums {
			if j-x.cost >= 0 && dp[j-x.cost] == dp[j]-1 {
				ans += strconv.Itoa(x.num)
				j -= x.cost
				break
			}
		}
	}

	fmt.Println(ans)
}

// sort ------------------------------------------------------------

type matchNum struct {
	num  int
	cost int
}

type sortArray []matchNum

func (s sortArray) Len() int           { return len(s) }
func (s sortArray) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortArray) Less(i, j int) bool { return s[i].num < s[j].num }

// -----------------------------------------------------------------

var (
	readString func() string
	readBytes  func() []byte
)

func init() {
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
