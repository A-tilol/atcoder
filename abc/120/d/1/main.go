package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Edge struct {
	a int
	b int
}

type Node struct {
	id       int
	parent   int
	children []int
}

type UnionFind struct {
	nodes []Node
}

func (u *UnionFind) init(n int) {
	nodes := make([]Node, n)
	for i := range nodes {
		nodes[i] = Node{
			id:       i,
			parent:   -1,
			children: make([]int, 0),
		}
	}
	u.nodes = nodes
}

func (u *UnionFind) getSize(i int) int {
	return 1 + len(u.nodes[i].children)
}

func (u *UnionFind) findRoot(i int) Node {
	parent := u.nodes[i].parent
	if parent == -1 {
		return u.nodes[i]
	}
	return u.nodes[parent]
}

func (u *UnionFind) unite(root, child Node) {
	if len(root.children) < len(child.children) {
		root, child = child, root
	}

	root.children = append(root.children, child.children...)
	root.children = append(root.children, child.id)
	u.nodes[root.id] = root

	u.nodes[child.id] = Node{
		id:     child.id,
		parent: root.id,
	}

	for _, id := range child.children {
		u.nodes[id].parent = root.id
	}
}

func main() {
	n, m := readInt(), readInt()

	edges := make([]Edge, m)
	for i := range edges {
		a, b := readInt()-1, readInt()-1
		edges[i] = Edge{a, b}
	}

	uf := &UnionFind{}
	uf.init(n)

	ans := make([]int, m)
	ans[m-1] = n * (n - 1) / 2

	for i := m - 1; i > 0; i-- {
		root1 := uf.findRoot(edges[i].a)
		root2 := uf.findRoot(edges[i].b)

		if root1.id == root2.id {
			ans[i-1] = ans[i]
			continue
		}
		ans[i-1] = ans[i] - uf.getSize(root1.id)*uf.getSize(root2.id)

		uf.unite(root1, root2)
	}

	for i := range ans {
		fmt.Println(ans[i])
	}
}

// sort ------------------------------------------------------------

type xxx struct {
	x int
}

type sortArray []xxx

func (s sortArray) Len() int           { return len(s) }
func (s sortArray) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortArray) Less(i, j int) bool { return s[i].x < s[j].x }

// -----------------------------------------------------------------

var (
	readString func() string
	readBytes  func() []byte
)

func init() {
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

func readInt() int {
	return int(readInt64())
}

func readInt64() int64 {
	i, err := strconv.ParseInt(readString(), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	return i
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
