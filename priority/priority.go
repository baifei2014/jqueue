// @Author jianglonghao
// @Date 2023/10/22
// @Description

// This example demonstrates a priority queue built using the heap interface.
package priority

import (
	"container/heap"
)

// An Item is something we manage in a priority queue.
type _item struct {
	value    any // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type _elements []*_item

type priority_queue struct {
	elements *_elements
	compare  _compare
}

type _compare func(i, j int) bool

var Greater _compare = func(i, j int) bool {
	if i > j {
		return true
	}
	return false
}

var Lesser _compare = func(i, j int) bool {
	if i < j {
		return true
	}
	return false
}

type PriorityQueue struct {
	queue *priority_queue
}

func New(compare _compare) PriorityQueue {

	pq := PriorityQueue{
		queue: &priority_queue{
			elements: &_elements{},
			compare:  compare,
		},
	}

	return pq
}

func (pq PriorityQueue) Empty() bool {
	if pq.queue.Len() > 0 {
		return false
	}
	return true
}

func (pq PriorityQueue) Put(val any, priority int) {
	item := &_item{
		value:    val,
		priority: priority,
	}
	heap.Push(pq.queue, item)
	pq.queue.update(item, val, priority)
}

func (pq PriorityQueue) Get() any {

	item := heap.Pop(pq.queue).(*_item)

	return item.value
}

func (es priority_queue) Len() int { return len(*es.elements) }

func (es priority_queue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return es.compare((*es.elements)[i].priority, (*es.elements)[j].priority)
}

func (es priority_queue) Swap(i, j int) {
	(*es.elements)[i], (*es.elements)[j] = (*es.elements)[j], (*es.elements)[i]
	(*es.elements)[i].index = i
	(*es.elements)[j].index = j
}

func (es *priority_queue) Push(x any) {
	n := len(*es.elements)
	item := x.(*_item)
	item.index = n
	*es.elements = append(*es.elements, item)
}

func (es *priority_queue) Pop() any {
	old := *es
	n := len(*old.elements)
	item := (*old.elements)[n-1]
	(*old.elements)[n-1] = nil // avoid memory leak
	item.index = -1            // for safety
	*es.elements = (*old.elements)[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *priority_queue) update(item *_item, value any, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
