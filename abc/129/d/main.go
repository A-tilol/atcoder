package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type row struct {
	s   int
	e   int
	len int
	idx int
}

type rowArr []row

func (a rowArr) Len() int           { return len(a) }
func (a rowArr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a rowArr) Less(i, j int) bool { return a[i].len < a[j].len }

func main1() {
	arr := nextLineToIntArray()
	h := arr[0]
	w := arr[1]
	S := readLines(h)

	x := byte('#')

	rows := make(rowArr, 0, h)
	for i, line := range S {
		s := 0
		for j := 0; j <= w; j++ {
			if j == w || line[j] == x {
				if j-s > 0 {
					r := row{
						s:   s,
						e:   j - 1,
						len: j - s,
						idx: i,
					}
					rows = append(rows, r)
					// fmt.Printf("%+v\n", r)
				}
				s = j + 1
			}
		}
	}
	// fmt.Println()

	cols := make(rowArr, 0, w)
	for i := 0; i < w; i++ {
		s := 0
		for j := 0; j <= h; j++ {
			if j == h || S[j][i] == x {
				if j-s > 0 {
					r := row{
						s:   s,
						e:   j - 1,
						len: j - s,
						idx: i,
					}
					cols = append(cols, r)
					// fmt.Printf("%+v\n", r)
				}
				s = j + 1
			}
		}
	}

	sort.Sort(sort.Reverse(rows))
	sort.Sort(sort.Reverse(cols))
	// fmt.Printf("%+v\n", rows)

	ans := 0
	for _, r := range rows {
		for _, c := range cols {
			if r.len+c.len-1 <= ans {
				break
			}
			if r.idx >= c.s && r.idx <= c.e {
				ans = r.len + c.len - 1
				break
			}
		}
	}

	fmt.Println(ans)
}

func main() {
	arr := nextLineToIntArray()
	h := arr[0]
	w := arr[1]
	S := readLines(h)

	x := byte('#')

	l := make([][]int, h)
	r := make([][]int, h)
	u := make([][]int, h)
	d := make([][]int, h)

	for i, line := range S {
		l[i] = make([]int, w)
		r[i] = make([]int, w)
		u[i] = make([]int, w)
		d[i] = make([]int, w)

		cnt := 0
		for j := 0; j < len(line); j++ {
			if line[j] == x {
				cnt = 0
			} else {
				l[i][j] = cnt
				cnt++
			}
		}
		cnt = 0
		for j := len(line) - 1; j >= 0; j-- {
			if line[j] == x {
				cnt = 0
			} else {
				r[i][j] = cnt
				cnt++
			}
		}
	}

	for i := 0; i < w; i++ {
		cnt := 0
		for j := 0; j < h; j++ {
			if S[j][i] == x {
				cnt = 0
			} else {
				u[j][i] = cnt
				cnt++
			}
		}
		cnt = 0
		for j := h - 1; j >= 0; j-- {
			if S[j][i] == x {
				cnt = 0
			} else {
				d[j][i] = cnt
				cnt++
			}
		}
	}
	// fmt.Printf("%+v\n", r)
	// fmt.Printf("%+v\n", l)
	// fmt.Printf("%+v\n", u)
	// fmt.Printf("%+v\n", d)

	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ans = max(ans, r[i][j]+l[i][j]+u[i][j]+d[i][j])
		}
	}
	fmt.Println(ans + 1)
}
