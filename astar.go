package astar

type Cost int

type Node interface {
	Nbs() []Node
	DistanceTo(other Node) Cost
	EstimateTo(other Node) Cost
}

func Astar(node0, goal Node) (path []Node) {
	var (
		closedset = map[Node]struct{}{}
		parents   = map[Node]Node{}
		g         = map[Node]Cost{node0: 0}
		f0        = node0.EstimateTo(goal)
		openq     = new(QMap)
	)
	openq.Init()
	openq.Add(node0, f0)

	for openq.Len() > 0 {
		x := openq.Pop()
		if x == goal {
			consPath(goal, parents, &path)
			return path
		}

		closedset[x] = struct{}{}

		for _, y := range x.Nbs() {
			if _, closed := closedset[y]; closed {
				continue
			}

			gEstimated := g[x] + x.DistanceTo(y)

			yQueued, updating := openq.Get(y)

			if updating {
				if gEstimated >= g[y] {
					continue
				}
			}

			parents[y] = x
			g[y] = gEstimated
			fy := gEstimated + y.EstimateTo(goal)

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
		if !ok {
			break
		}
		*path = append(*path, node)
		node = parent
	}
	reverse(*path)
}

func reverse(path []Node) {
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
}
