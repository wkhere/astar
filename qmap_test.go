package astar

import (
	"fmt"
	"testing"
)

// dummy Node impl
type node string

func (n node) Nbs() []Node             { return nil }
func (n node) DistanceTo(n2 Node) Cost { return 1 }
func (n node) EstimateTo(n2 Node) Cost { return 10 }

func qmapExample() []Node {
	q := new(QMap)
	q.Init()
	q.Add(node("foo"), 10)
	q.Add(node("five"), 5)
	q.Add(node("five2"), 5)
	item, _ := q.Get(node("foo"))
	q.Update(item, 3)
	q.Add(node("two"), 2)
	res := make([]Node, 0, 4)
	for q.Len() > 0 {
		v := q.Pop()
		res = append(res, v)
	}
	return res
}

func ExampleQMap() {
	fmt.Print(qmapExample())
	// Output: [two foo five five2]
}

func BenchmarkQMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		qmapExample()
	}
}
