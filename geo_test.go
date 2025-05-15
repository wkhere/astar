package astar

import (
	"fmt"
	"testing"

	"github.com/wkhere/astar/graphs/geo"
)

var Point = geo.Point

// TODO:
// need more dense graph to test & benchmark non-trivial paths

func ExampleL() {
	fmt.Println(Astar(L("Wałcz"), L("")))
	fmt.Println(Astar(L("Wałcz"), L("Wałcz")))
	fmt.Println(Astar(L("Wałcz"), L("Warszawa")))
	fmt.Println(Astar(L("Warszawa"), L("Wałcz")))
	fmt.Println(Astar(L("Wałcz"), L("Poznań")))
	fmt.Println(Astar(L("Wałcz"), L("Cieszyn")))
	fmt.Println(Astar(L("Cieszyn"), L("Wałcz")))
	// Output:
	// []
	// []
	// [Warszawa]
	// [Wałcz]
	// [Trzcianka Poznań]
	// [Trzcianka Poznań Leszno Katowice Cieszyn]
	// [Katowice Leszno Poznań Trzcianka Wałcz]
}

func BenchmarkGeo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Astar(L("Wałcz"), L("Wałcz"))
		Astar(L("Wałcz"), L("Warszawa"))
		Astar(L("Wałcz"), L("Poznań"))
	}
}

type L string // Location

func (l1 L) Nbs() []Node {
	return nbs[l1]
}

func (l1 L) DistanceTo(n2 Node) (v Cost) {
	l2 := n2.(L)
	v, ok := distLookup(l1, l2)
	if !ok {
		panic(fmt.Sprintf("no dist for %v,%v", l1, l2))
	}
	return
}

func (l1 L) EstimateTo(n2 Node) Cost {
	return Cost(geo.H(coords[l1], coords[n2.(L)]))
}

var coords = map[L]geo.Pt{
	"Cieszyn":   Point(49.75, 63.19),
	"Leszno":    Point(51.8454, 16.5748),
	"Wałcz":     Point(53.283853, 16.470173),
	"Trzcianka": Point(53.0427712, 16.3763841),
	"Piła":      Point(53.1347933, 16.6195561),
	"Poznań":    Point(52.408031, 16.920613),
	"Warszawa":  Point(52.230069, 21.018513),
	"Katowice":  Point(50.26, 19.02),
}

type locationPair struct{ l1, l2 L }

var distances = map[locationPair]Cost{
	// these are arbitrary distances taken from real maps
	{"Wałcz", "Trzcianka"}:   31,
	{"Trzcianka", "Poznań"}:  88,
	{"Wałcz", "Piła"}:        28,
	{"Piła", "Poznań"}:       96,
	{"Wałcz", "Warszawa"}:    421,
	{"Poznań", "Warszawa"}:   310,
	{"Cieszyn", "Katowice"}:  74,
	{"Poznań", "Leszno"}:     81,
	{"Leszno", "Katowice"}:   298,
	{"Warszawa", "Katowice"}: 295,
}

var nbs = map[Node][]Node{}

func init() {
	for k := range distances {
		nbs[k.l1] = append(nbs[k.l1], k.l2)
		nbs[k.l2] = append(nbs[k.l2], k.l1)
	}
}

func distLookup(l1, l2 L) (v Cost, ok bool) {
	v, ok = distances[locationPair{l1, l2}]
	if ok {
		return
	}
	v, ok = distances[locationPair{l2, l1}]
	return
}
