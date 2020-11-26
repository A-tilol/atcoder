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
	h, w, n, m := readInt(), readInt(), readInt(), readInt()
	grid := make([][]int, h)
	vgrid := make([][]int, h)
	hgrid := make([][]int, h)
	for i := range grid {
		grid[i] = make([]int, w)
		vgrid[i] = make([]int, w)
		hgrid[i] = make([]int, w)
	}

	for i := 0; i < n; i++ {
		a, b := readInt()-1, readInt()-1
		grid[a][b] = 1
	}

	for i := 0; i < m; i++ {
		a, b := readInt()-1, readInt()-1
		grid[a][b] = -1
	}

	// deb
	// for i := range grid {
	// 	fmt.Println(grid[i])
	// }
	// fmt.Println()

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] != 1 {
				continue
			}
			hgrid[i][j] = 1

			//left
			jj := j - 1
			for jj >= 0 && grid[i][jj] == 0 {
				hgrid[i][jj] = 1
				jj--
			}
			//right
			j++
			for j < w && grid[i][j] != -1 {
				hgrid[i][j] = 1
				j++
			}
		}
	}

	for j := 0; j < w; j++ {
		for i := 0; i < h; i++ {
			if grid[i][j] != 1 {
				continue
			}
			vgrid[i][j] = 1

			//up
			ii := i - 1
			for ii >= 0 && grid[ii][j] == 0 {
				vgrid[ii][j] = 1
				ii--
			}
			//down
			i++
			for i < h && grid[i][j] != -1 {
				vgrid[i][j] = 1
				i++
			}
		}
	}

	// fmt.Println("h grid")
	// for i := 0; i < h; i++ {
	// 	fmt.Println(hgrid[i])
	// }

	// fmt.Println("v grid")
	// for i := 0; i < h; i++ {
	// 	fmt.Println(vgrid[i])
	// }

	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if hgrid[i][j] == 1 || vgrid[i][j] == 1 {
				ans++
			}
		}
	}
	fmt.Println(ans)
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
