package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

var r90 = [][]int{
	{0, -1, 0},
	{1, 0, 0},
	{0, 0, 1},
}
var rn90 = [][]int{
	{0, 1, 0},
	{-1, 0, 0},
	{0, 0, 1},
}
var symmetryXp = [][]int{
	{-1, 0, 0},
	{0, 1, 0},
	{0, 0, 1},
}
var symmetryYp = [][]int{
	{1, 0, 0},
	{0, -1, 0},
	{0, 0, 1},
}

func product(a [][]int, b [][]int) [][]int {
	c := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	for i := range b {
		for j := range b[i] {
			for k := range a[i] {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}

func main() {
	n := readInt()
	xys := make([][][]int, n)
	for i := range xys {
		xys[i] = [][]int{
			{readInt()},
			{readInt()},
			{1},
		}
	}
	m := readInt()
	ops := make([][]int, m)
	for i := range ops {
		o := readInt()
		p := 0
		if o >= 3 {
			p = readInt()
		}
		ops[i] = []int{o, p}
	}
	ops = append(ops, []int{0, 0})
	q := readInt()
	nabs := make(map[int][][]int)
	for i := 0; i < q; i++ {
		a, b := readInt(), readInt()-1
		if _, ok := nabs[a]; !ok {
			nabs[a] = [][]int{}
		}
		nabs[a] = append(nabs[a], []int{i, b})
	}

	trans := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	ansxys := make([][]int, q)
	for i := range ops {
		for _, nab := range nabs[i] {
			p := product(trans, xys[nab[1]])
			ansxys[nab[0]] = []int{p[0][0], p[1][0]}
		}
		if ops[i][0] == 1 {
			trans = product(rn90, trans)
		} else if ops[i][0] == 2 {
			trans = product(r90, trans)
		} else if ops[i][0] == 3 {
			symmetryXp[0][2] = 2 * ops[i][1]
			trans = product(symmetryXp, trans)
		} else if ops[i][0] == 4 {
			symmetryYp[1][2] = 2 * ops[i][1]
			trans = product(symmetryYp, trans)
		}
	}

	for i := range ansxys {
		fmt.Println(ansxys[i][0], ansxys[i][1])
	}

}

func main1() {
	n := readInt()
	xys := make([][][]int, n)
	for i := range xys {
		xys[i] = [][]int{
			{readInt()},
			{readInt()},
			{1},
		}
	}
	m := readInt()
	ops := make([][]int, m)
	for i := range ops {
		o := readInt()
		p := 0
		if o >= 3 {
			p = readInt()
		}
		ops[i] = []int{o, p}
	}
	ops = append(ops, []int{0, 0})
	q := readInt()
	nabs := make([][]int, q)
	for i := range nabs {
		nabs[i] = []int{i, readInt(), readInt() - 1}
	}

	// log.Println("xys", xys)
	// log.Println("ops", ops)
	// log.Println("nabs", nabs)

	sort.Slice(nabs, func(i, j int) bool {
		return nabs[i][1] < nabs[j][1]
	})
	// log.Println("nabs", nabs)

	k := 0
	trans := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	ansxys := make([][]int, q)
	for i := range ops {
		for k < q && nabs[k][1] == i {
			p := product(trans, xys[nabs[k][2]])
			ansxys[nabs[k][0]] = []int{p[0][0], p[1][0]}
			k++
		}
		if ops[i][0] == 1 {
			trans = product(rn90, trans)
		} else if ops[i][0] == 2 {
			trans = product(r90, trans)
		} else if ops[i][0] == 3 {
			symmetryXp[0][2] = 2 * ops[i][1]
			trans = product(symmetryXp, trans)
		} else if ops[i][0] == 4 {
			symmetryYp[1][2] = 2 * ops[i][1]
			trans = product(symmetryYp, trans)
		}
	}

	// log.Println(ansxys)

	for i := range ansxys {
		fmt.Println(ansxys[i][0], ansxys[i][1])
	}

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

func readInts(n int) []int {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = readInt()
	}
	return arr
}

func readFloat64() float64 {
	f, err := strconv.ParseFloat(readString(), 64)
	if err != nil {
		panic(err.Error())
	}
	return f
}

func readFloat64s(n int) []float64 {
	arr := make([]float64, n)
	for i := range arr {
		arr[i] = readFloat64()
	}
	return arr
}

func min(a ...int) int {
	x := math.MaxInt64
	for i := range a {
		if x > a[i] {
			x = a[i]
		}
	}
	return x
}

func max(a ...int) int {
	x := math.MinInt64
	for i := range a {
		if x < a[i] {
			x = a[i]
		}
	}
	return x
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
