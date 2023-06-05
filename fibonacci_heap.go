package main

import (
	"fmt"
)

type HeapTree struct {
	Value    *int
	Children []*HeapTree
	Order    *int
}

func (h *HeapTree) AddAtEnd(ele *HeapTree) {
	h.Children = append(h.Children, ele)
	*h.Order += 1
}

// definition
type FibonacciHeap struct {
	Trees []*HeapTree
	Least *HeapTree
	Count *int
}

func (f *FibonacciHeap) Insert(el *HeapTree) {
	f.Trees = InsertEle(f.Trees, el)
	if *f.Count == 0 || *el.Value <= *f.Least.Value {
		f.Least = el
	}
	*f.Count += 1
}

func (f *FibonacciHeap) GetMin() HeapTree {
	return *f.Least
}

func (f *FibonacciHeap) ExtractMin() (HeapTree, error) {

	// nil Heap
	if *f.Count == 0 {
		return HeapTree{}, fmt.Errorf("nil heap")
	}

	rt := f.Least
	f.Trees = RemoveEle(f.Trees, f.Least)
	*f.Count -= 1

	if *f.Count > 0 {
		f.Least = f.Trees[0]
	} else {
		f.Least = &HeapTree{}
	}

	for _, ch := range rt.Children {
		f.Insert(ch)
	}

	for _, tr := range f.Trees {
		if *f.Count == 0 || *f.Least.Value > *tr.Value {
			f.Least = tr
		}
	}
	if *f.Count != 0 {
		f.Heapify()
	}
	return *rt, nil
}

func (f *FibonacciHeap) Heapify() {
	Orders := make(map[int]*HeapTree)
	for *f.Count > 0 {
		x := f.Trees[0]
		f.Trees = RemoveEle(f.Trees, x)
		*f.Count -= 1
		order := *x.Order
		for y, ok := Orders[order]; ok; y, ok = Orders[order] {
			delete(Orders, order)
			if *y.Value < *x.Value {
				y, x = x, y
			}
			x.AddAtEnd(y)
			order += 1
		}
		Orders[order] = x
	}
	for _, tr := range Orders {
		f.Insert(tr)
	}

}
