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
	n, k := readInt(), readInt()
	p := make([]int, n)
	c := make([]int, n)
	for i := range p {
		p[i] = readInt() - 1
	}
	for i := range c {
		c[i] = readInt()
	}

	ans := math.MinInt64
	for i := range p {
		memo := make([]int, n)
		for j := range memo {
			memo[j] = -1
		}

		cur, cnt := i, 0
		for memo[cur] == -1 {
			memo[cur] = cnt
			cur = p[cur]
			cnt++
		}
		loopStartP := cur
		cntOneLoopEnd := cnt

		cur = p[loopStartP]
		scoreOfOneLoop, lenOfOneLoop := c[loopStartP], 1
		for cur != loopStartP {
			lenOfOneLoop++
			scoreOfOneLoop += c[cur]
			cur = p[cur]
		}

		// sum pi until first loop end
		tmpAns := 0
		cur, cnt = i, 0
		loopStartPCnt := 0
		for cnt < k {
			if cur == loopStartP {
				loopStartPCnt++
				if loopStartPCnt == 2 {
					break
				}
			}
			cur = p[cur]
			tmpAns += c[cur]
			ans = max(ans, tmpAns)
			cnt++
		}

		if cntOneLoopEnd > k || scoreOfOneLoop <= 0 {
			continue
		}

		cur = i
		tmpK := k
		scoreUntilLoopStart := 0
		for cur != loopStartP {
			cur = p[i]
			tmpAns += c[cur]
			tmpK--
		}

		tmpAns = scoreUntilLoopStart
		tmpAns += (tmpK/lenOfOneLoop - 1) * scoreOfOneLoop
		tmpK = tmpK%lenOfOneLoop + lenOfOneLoop

		cur = i
		for j := 0; j < tmpK; j++ {
			cur = p[cur]
			tmpAns += c[cur]
			ans = max(ans, tmpAns)
		}
	}

	fmt.Println(ans)
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
