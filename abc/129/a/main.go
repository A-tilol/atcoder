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

func submit(line string) int {
	parts := strings.Split(line, " ")
	p, _ := strconv.Atoi(parts[0])
	q, _ := strconv.Atoi(parts[1])
	r, _ := strconv.Atoi(parts[2])

	return min(p+q, min(q+r, r+p))
}

func main() {
	line := nextLine()
	ans := submit(line)
	fmt.Println(ans)
}
