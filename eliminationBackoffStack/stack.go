package stack

import (
	"math/rand"
	"sync/atomic"
)

type Stack[T any] struct {
	head                    atomic.Value
	eliminationBackoffArray *EliminationBackoffArray[T]
}

type Node[T any] struct {
	element *T
	next    *Node[T]
}

type EliminationBackoffArray[T any] struct {
	exchangers []*Exchanger[T]
}

type Exchanger[T any] struct {
	item atomic.Value
}

type Item[T any] struct {
	state ItemState
	data  *T
}

type ItemState int

const (
	Empty ItemState = iota
	Waiting
	Busy
)

func newEliminationBackoffArray[T any](size int) *EliminationBackoffArray[T] {
	array := &EliminationBackoffArray[T]{
		exchangers: make([]*Exchanger[T], size),
	}

	for index := 0; index < size; index++ {
		array.exchangers[index] = &Exchanger[T]{}
		array.exchangers[index].item.Store(&Item[T]{state: Empty})
	}

	return array
}

func NewEBStack[T any](size int) *Stack[T] {
	stack := &Stack[T]{
		eliminationBackoffArray: newEliminationBackoffArray[T](size),
	}

	stack.head.Store((*Node[T])(nil))

	return stack
}

func newNode[T any](element *T) *Node[T] {
	return &Node[T]{
		element: element,
		next:    nil,
	}
}

func (s *Stack[T]) Push(element T) {
	node := newNode[T](&element)

	for {
		// try in main stack
		if s.tryPush(node) {
			return
		}

		// try in elimination array
		if s.tryEliminatePush(node) {
			return
		}
	}
}

func (s *Stack[T]) tryPush(node *Node[T]) bool {
	currentHead := s.head.Load().(*Node[T])
	node.next = currentHead
	return s.head.CompareAndSwap(currentHead, node)
}

func (s *Stack[T]) tryEliminatePush(node *Node[T]) bool {
	slot := rand.Intn(len(s.eliminationBackoffArray.exchangers))
	exchanger := s.eliminationBackoffArray.exchangers[slot]

	currentItem := exchanger.item.Load().(*Item[T])

	// skip if in Waiting or Busy state
	if currentItem.state == Waiting || currentItem.state == Busy {
		return false
	}

	newItem := &Item[T]{state: Waiting, data: node.element}

	// check if currentItem is valid and exchange with newItem
	return exchanger.item.CompareAndSwap(currentItem, newItem)
}

func (s *Stack[T]) Pop() *T {
	for {
		// try in main stack
		if val, ok := s.tryPop(); ok {
			return val
		}

		// try in elimination array
		if val, ok := s.tryEliminatePop(); ok {
			return val
		}
	}
}

func (s *Stack[T]) tryPop() (*T, bool) {
	currentHead := s.head.Load().(*Node[T])
	if currentHead == nil {
		return nil, true
	}

	if s.head.CompareAndSwap(currentHead, currentHead.next) {
		return currentHead.element, true
	}

	return nil, false
}

func (s *Stack[T]) tryEliminatePop() (*T, bool) {
	slot := rand.Intn(len(s.eliminationBackoffArray.exchangers))
	exchanger := s.eliminationBackoffArray.exchangers[slot]

	currentItem := exchanger.item.Load().(*Item[T])

	// skip if in Empty or Busy state
	if currentItem.state == Empty || currentItem.state == Busy {
		return nil, false
	}

	newItem := &Item[T]{state: Busy, data: currentItem.data}

	// check if currentItem is valid and exchange with newItem
	// and then reset the exhange
	if exchanger.item.CompareAndSwap(currentItem, newItem) {
		data := currentItem.data
		exchanger.item.Store(&Item[T]{state: Empty})
		return data, true
	}

	return nil, true
}

func (s *Stack[T]) IsEmpty() bool {
	return s.head.Load() == nil
}
