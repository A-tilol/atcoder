package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
)

var nums = [][]byte{
	[]byte{46, 35, 35, 35, 46, 35, 46, 35, 46, 35, 46, 35, 46, 35, 46, 35, 46, 35, 35, 35},
	[]byte{46, 46, 35, 46, 46, 35, 35, 46, 46, 46, 35, 46, 46, 46, 35, 46, 46, 35, 35, 35},
	[]byte{46, 35, 35, 35, 46, 46, 46, 35, 46, 35, 35, 35, 46, 35, 46, 46, 46, 35, 35, 35},
	[]byte{46, 35, 35, 35, 46, 46, 46, 35, 46, 35, 35, 35, 46, 46, 46, 35, 46, 35, 35, 35},
	[]byte{46, 35, 46, 35, 46, 35, 46, 35, 46, 35, 35, 35, 46, 46, 46, 35, 46, 46, 46, 35},
	[]byte{46, 35, 35, 35, 46, 35, 46, 46, 46, 35, 35, 35, 46, 46, 46, 35, 46, 35, 35, 35},
	[]byte{46, 35, 35, 35, 46, 35, 46, 46, 46, 35, 35, 35, 46, 35, 46, 35, 46, 35, 35, 35},
	[]byte{46, 35, 35, 35, 46, 46, 46, 35, 46, 46, 46, 35, 46, 46, 46, 35, 46, 46, 46, 35},
	[]byte{46, 35, 35, 35, 46, 35, 46, 35, 46, 35, 35, 35, 46, 35, 46, 35, 46, 35, 35, 35},
	[]byte{46, 35, 35, 35, 46, 35, 46, 35, 46, 35, 35, 35, 46, 46, 46, 35, 46, 35, 35, 35},
}

func main() {
	n := readInt()
	s := make([][]byte, 5)
	for i := range s {
		s[i] = readBytes()
	}

	for i := 0; i < n; i++ {
		num := make([]byte, 0, 20)
		for j := 0; j < 5; j++ {
			num = append(num, s[j][i*4:(i+1)*4]...)
		}

		for k := range nums {
			if reflect.DeepEqual(num, nums[k]) {
				fmt.Print(k)
				break
			}
		}
	}
}

func createNums() {
	n := readInt()
	s := make([][]byte, 5)
	for i := range s {
		s[i] = readBytes()
	}

	nums := make([][]byte, 10)
	for i := range nums {
		for j := 0; j < 5; j++ {
			nums[i] = append(nums[i], s[j][i*4:(i+1)*4]...)
		}
	}

	for i := range nums {
		fmt.Println(nums[i])
	}
	fmt.Println(n)
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
