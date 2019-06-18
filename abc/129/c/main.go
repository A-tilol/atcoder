package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextLineToStringArray() []string {
	line := nextLine()
	return strings.Split(line, " ")
}

func nextLineToIntArray() []int {
	parts := nextLineToStringArray()
	arr := make([]int, len(parts))
	for i, s := range parts {
		arr[i], _ = strconv.Atoi(s)
	}
	return arr
}

func nextLineToInt64Array() []int64 {
	parts := nextLineToStringArray()
	arr := make([]int64, len(parts))
	for i, s := range parts {
		arr[i], _ = strconv.ParseInt(s, 10, 64)
	}
	return arr
}

func nextLineToFloat32Array() []float32 {
	parts := nextLineToStringArray()
	arr := make([]float32, len(parts))
	for i, s := range parts {
		n, _ := strconv.ParseFloat(s, 32)
		arr[i] = float32(n)
	}
	return arr
}

func nextLineToFloat64Array() []float64 {
	parts := nextLineToStringArray()
	arr := make([]float64, len(parts))
	for i, s := range parts {
		arr[i], _ = strconv.ParseFloat(s, 64)
	}
	return arr
}

func readLines(n int) []string {
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		arr[i] = nextLine()
	}
	return arr
}

func readLinesInt(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.Atoi(nextLine())
	}
	return arr
}

func readLinesInt64(n int) []int64 {
	arr := make([]int64, n)
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.ParseInt(nextLine(), 10, 64)
	}
	return arr
}

func readLinesFloat32(n int) []float32 {
	arr := make([]float32, n)
	for i := 0; i < n; i++ {
		n, _ := strconv.ParseFloat(nextLine(), 32)
		arr[i] = float32(n)
	}
	return arr
}

func readLinesFloat64(n int) []float64 {
	arr := make([]float64, n)
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.ParseFloat(nextLine(), 64)
	}
	return arr
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
	if a > 0 {
		return a
	}
	return -a
}

func submit(n int, line string) int {
	return 0
}

var memo = make(map[int64]int64)

func fib(i int64) int64 {
	if v, ok := memo[i]; ok {
		return v
	}
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	memo[i] = (fib(i-2) + fib(i-1)) % 1000000007
	return memo[i]
}

func main() {
	arr := nextLineToIntArray()
	n := arr[0]
	m := arr[1]
	A := readLinesInt64(m)
	A = append(A, int64(n+1))

	var i int64
	var ans int64 = 1
	for _, a := range A {
		t := a - i
		if t == 0 {
			ans = 0
			break
		}
		ans = (ans * fib(t)) % 1000000007
		i = a + 1
	}
	fmt.Println(ans)
}
