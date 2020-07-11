package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	n := readInt()
	s := readString()

	bitPlusCnt := strings.Count(s, "1") + 1
	bitMinusCnt := bitPlusCnt - 2
	if bitMinusCnt == 0 {
		bitMinusCnt = -1
	}
	// fmt.Println(bitCnt, bitMinusCnt, bitPlusCnt)

	plusR := make([]int, n)
	minusR := make([]int, n)
	plusR[n-1] = 1 % bitPlusCnt
	minusR[n-1] = 1 % bitMinusCnt
	for i := n - 2; i >= 0; i-- {
		plusR[i] = plusR[i+1] * 2 % bitPlusCnt
		minusR[i] = minusR[i+1] * 2 % bitMinusCnt
	}
	// fmt.Println(plusR)
	// fmt.Println(minusR)

	rPlus := 0
	rMinus := 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == '1' {
			rPlus = (rPlus + plusR[i]) % bitPlusCnt
			rMinus = (rMinus + minusR[i]) % bitMinusCnt
		}
	}

	for i := 0; i < n; i++ {
		r := 0
		if s[i] == '1' {
			if bitMinusCnt == -1 {
				fmt.Println(0)
				continue
			}
			r = ((rMinus-minusR[i])%bitMinusCnt + bitMinusCnt) % bitMinusCnt
		} else {
			r = (rPlus + plusR[i]) % bitPlusCnt
		}

		ans := 1
		for r != 0 {
			bcnt := strings.Count(fmt.Sprintf("%b", r), "1")
			r = r % bcnt
			ans++
		}

		fmt.Println(ans)
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
