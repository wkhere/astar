package astar

import (
	"fmt"
	"testing"
)

func openqExample() []Node {
	openq := new(OpenQS)
	openq.Init()
	openq.Add("foo", 10)
	openq.Add("five", 5)
	openq.Add("five2", 5)
	item, _ := openq.Item("foo")
	openq.Update(item, 3)
	openq.Add("two", 2)
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
