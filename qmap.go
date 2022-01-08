package astar

import "container/heap"

type qitem struct {
	value    Node
	priority Cost
	index    int
	t        int // 'timestamp' so that items added earlier have it lower
}
type queue []*qitem

// QMap is a hybrid of priority queue and map,
// allowing for O(log(N)) Add/Update and O(1) Pop/Get.
type QMap struct {
	q        queue
	m        map[Node]*qitem
	tcounter int
}

// API

func (qm *QMap) Init() {
	qm.q = make(queue, 0, 10)
	qm.m = make(map[Node]*qitem)
}

func (qm *QMap) Add(v Node, priority Cost) {
	qm.tcounter++
	x := &qitem{v, priority, -1, qm.tcounter}
	heap.Push(&qm.q, x)
	qm.m[v] = x
}

// Update can only work on the element obtained by Get.
func (qm *QMap) Update(x *qitem, newPriority Cost) {
	qm.tcounter++
	x.t = qm.tcounter
	x.priority = newPriority
	heap.Fix(&qm.q, x.index)
}

func (qm *QMap) Pop() Node {
	v := heap.Pop(&qm.q).(*qitem).value
	delete(qm.m, v)
	return v
}

// Get returns element which can be updated
// by Update, plus a bool indicating if it existed.
func (qm *QMap) Get(v Node) (*qitem, bool) {
	x, ok := qm.m[v]
	return x, ok
}

func (qm *QMap) Len() int { return len(qm.m) }

// heap.Interface impl

func (q queue) Len() int { return len(q) }

func (q queue) Less(i, j int) bool {
	pri1, pri2 := q[i].priority, q[j].priority
	if pri1 == pri2 {
		return q[i].t < q[j].t
	}
	return pri1 < pri2
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *queue) Push(obj interface{}) {
	x := obj.(*qitem)
	x.index = len(*q)
	*q = append(*q, x)
}

func (q *queue) Pop() interface{} {
	last := len(*q) - 1
	x := (*q)[last]
	*q = (*q)[0:last]
	return x
}
