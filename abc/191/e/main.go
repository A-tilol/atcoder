package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"runtime"
	"strconv"
)

// Item is the type we manage in the priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push is used to add a new element to the heap.
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop removes the element with the highest priority from the heap.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type pair struct {
	l int
	c int
}

func main() {
	debug.logging = false

	n, m := readInt(), readInt()
	lmap := make(map[int]map[int]int)
	for i := 0; i < m; i++ {
		a, b, c := readInt()-1, readInt()-1, readInt()
		if _, ok := lmap[a]; !ok {
			lmap[a] = make(map[int]int)
		}
		if _, ok := lmap[a][b]; !ok {
			lmap[a][b] = c
		} else {
			lmap[a][b] = min(lmap[a][b], c)
		}
	}

	for i := 0; i < n; i++ {
		pq := PriorityQueue{{value: i, priority: 0}}
		heap.Init(&pq)
		cmap := make([]int, n)
		for k := range cmap {
			cmap[k] = math.MaxInt64
		}
		for len(pq) > 0 {
			cur := heap.Pop(&pq).(*Item)
			debug.println(cur.value, cur.priority, lmap[cur.value])
			if cmap[cur.value] <= cur.priority {
				continue
			}
			if cur.priority != 0 {
				cmap[cur.value] = cur.priority
			}
			for next, c := range lmap[cur.value] {
				if cmap[next] <= cur.priority {
					continue
				}
				heap.Push(&pq, &Item{value: next, priority: cur.priority + c})
			}
		}
		if cmap[i] == math.MaxInt64 {
			fmt.Println(-1)
		} else {
			fmt.Println(cmap[i])
		}
		debug.println(cmap)
	}
}

var readString func() string
var debug *debugLogger

func init() {
	log.SetFlags(log.Lshortfile)
	readString = newReadString(os.Stdin)
	debug = newDebugLogger(true)
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

type debugLogger struct {
	logging bool
}

func newDebugLogger(logging bool) *debugLogger {
	return &debugLogger{logging: logging}
}

func (d *debugLogger) print(args ...interface{}) {
	if d.logging {
		fmt.Print(args...)
	}
}

func (d *debugLogger) printf(format string, args ...interface{}) {
	if d.logging {
		_, _, line, _ := runtime.Caller(1)
		format = fmt.Sprintf("(l%d) ", line) + format
		fmt.Printf(format, args...)
	}
}

func (d *debugLogger) println(args ...interface{}) {
	if d.logging {
		_, _, line, _ := runtime.Caller(1)
		format := fmt.Sprintf("(l%d)", line)
		args = append([]interface{}{format}, args...)
		fmt.Println(args...)
	}
}
