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
	s := readBytes()

	multiples := make([][]int, 0)
	num := 8
	for {
		numStr := strconv.Itoa(num)
		if len(numStr) > 3 {
			break
		}

		if num%10 == 0 {
			num += 8
			continue
		}

		multiple := make([]int, 10)
		multiple[0] = len(numStr)
		for i := range numStr {
			d := int(numStr[i]) - 48
			multiple[d]++
		}
		multiples = append(multiples, multiple)
		num += 8
	}

	dToCnt := map[int]int{}
	for _, b := range s {
		d := int(b) - 48
		dToCnt[d]++
	}

	for _, mul := range multiples {
		flg := true
		if len(s) == 1 && mul[0] != 1 {
			continue
		}
		if len(s) == 2 && mul[0] != 2 {
			continue
		}
		if len(s) >= 3 && mul[0] != 3 {
			continue
		}

		for i := 1; i < 10; i++ {
			if dToCnt[i] < mul[i] {
				flg = false
				break
			}
		}
		if flg {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
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
