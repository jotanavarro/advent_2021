package dataStructures

import "container/heap"

type Item struct {
	Value    rune
	Priority int
	index    int
}

type RunePriorityQueue []*Item

func (rpq RunePriorityQueue) Len() int           { return len(rpq) }
func (rpq RunePriorityQueue) Less(i, j int) bool { return rpq[i].Priority > rpq[j].Priority }
func (rpq RunePriorityQueue) Swap(i, j int) {
	rpq[i], rpq[j] = rpq[j], rpq[i]
	rpq[i].index = i
	rpq[j].index = j
}

func (rpq *RunePriorityQueue) Push(x interface{}) {
	n := len(*rpq)
	item := x.(*Item)
	item.index = n
	*rpq = append(*rpq, item)
}

func (rpq *RunePriorityQueue) Pop() interface{} {
	old := *rpq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*rpq = old[0 : n-1]
	return item
}

func (rpq *RunePriorityQueue) update(item *Item, value rune, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(rpq, item.index)
}
