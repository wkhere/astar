package astar

import (
	"fmt"
	"testing"

	"github.com/wkhere/astar/graphs/geo"
)

// TODO:
// need more dense graph to test & benchmark non-trivial paths

func ExampleGeo() {
	g := Geo{}
	fmt.Println(Astar(g, "Wałcz", ""))
	fmt.Println(Astar(g, "Wałcz", "Wałcz"))
	fmt.Println(Astar(g, "Wałcz", "Warszawa"))
	fmt.Println(Astar(g, "Warszawa", "Wałcz"))
	fmt.Println(Astar(g, "Wałcz", "Poznań"))
	// Output:
	// []
	// []
	// [Warszawa]
	// [Wałcz]
	// [Trzcianka Poznań]
}

func BenchmarkGeo(b *testing.B) {
	g := Geo{}
	for n := 0; n < b.N; n++ {
		Astar(g, "Wałcz", "Wałcz")
		Astar(g, "Wałcz", "Warszawa")
		Astar(g, "Wałcz", "Poznań")
	}
}

type Geo struct{}

func (g Geo) Nbs(node Node) []Node {
	return nbs[node]
}

func (g Geo) Dist(n1, n2 Node) (v Cost) {
	v, ok := distLookup(n1, n2)
	if !ok {
		panic(fmt.Sprintf("no dist for %v,%v", n1, n2))
	}
	return
}

func (g Geo) H(n1, n2 Node) Cost {
	return Cost(geo.H(coords[n1], coords[n2]))
}

var coords = map[Node]geo.Pt{
	"Wałcz":     {53.283853, 16.470173},
	"Trzcianka": {53.0427712, 16.3763841},
	"Piła":      {53.1347933, 16.6195561},
	"Poznań":    {52.408031, 16.920613},
	"Warszawa":  {52.230069, 21.018513},
}

type nodePair struct{ n1, n2 Node }

var distances = map[nodePair]Cost{
	// these are arbitrary distances taken from real maps
	{"Wałcz", "Trzcianka"}:  31,
	{"Trzcianka", "Poznań"}: 88,
	{"Wałcz", "Piła"}:       28,
	{"Piła", "Poznań"}:      96,
	{"Wałcz", "Warszawa"}:   421,
	{"Poznań", "Warszawa"}:  310,
}

var nbs = map[Node][]Node{}

func init() {
	for k := range distances {
		nbs[k.n1] = append(nbs[k.n1], k.n2)
		nbs[k.n2] = append(nbs[k.n2], k.n1)
	}
}

func distLookup(n1, n2 Node) (v Cost, ok bool) {
	v, ok = distances[nodePair{n1, n2}]
	if ok {
		return
	}
	v, ok = distances[nodePair{n2, n1}]
	return
}
