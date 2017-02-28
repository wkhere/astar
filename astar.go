package astar

type Cost int
type Node interface{}

type Graph interface {
	Nbs(Node) []Node
	Dist(n1, n2 Node) Cost
	H(n1, n2 Node) Cost
}

func Astar(graph Graph, node0, goal Node) (path []Node) {
	closedset := map[Node]struct{}{}
	parents := map[Node]Node{}
	g := map[Node]Cost{node0: 0}
	f0 := graph.H(node0, goal)
	openq := new(OpenQS)
	openq.Init()
	openq.Add(node0, f0)

	for openq.Len() > 0 {
		x := openq.Pop()
		if x == goal {
			consPath(goal, parents, &path)
			return path
		}

		closedset[x] = struct{}{}

		for _, y := range graph.Nbs(x) {
			if _, closed := closedset[y]; closed {
				continue
			}

			gEstimated := g[x] + graph.Dist(x, y)

			yQueued, updating := openq.Item(y)

			if updating {
				if gEstimated >= g[y] {
					continue
				}
			}

			parents[y] = x
			g[y] = gEstimated
			fy := graph.H(y, goal) + gEstimated

			if updating {
				openq.Update(yQueued, fy)
			} else {
				openq.Add(y, fy)
			}
		}
	}
	return
}

func consPath(node Node, parents map[Node]Node, path *[]Node) {
	for {
		parent, ok := parents[node]
		if ok {
			*path = append(*path, node)
			node = parent
		} else {
			break
		}
	}
	reverse(*path)
}

func reverse(path []Node) {
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
}
