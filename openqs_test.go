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

func openqExample() []Node {
	openq := new(OpenQS)
	openq.Init()
	openq.Add(node("foo"), 10)
	openq.Add(node("five"), 5)
	openq.Add(node("five2"), 5)
	item, _ := openq.Item(node("foo"))
	openq.Update(item, 3)
	openq.Add(node("two"), 2)
	res := make([]Node, 0, 4)
	for openq.Len() > 0 {
		v := openq.Pop()
		res = append(res, v)
	}
	return res
}

func ExampleOpenQS() {
	fmt.Print(openqExample())
	// Output: [two foo five five2]
}

func BenchmarkOpenQS(b *testing.B) {
	for n := 0; n < b.N; n++ {
		openqExample()
	}
}
