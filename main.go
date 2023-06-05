package main

import (
	"fmt"

	fake "github.com/brianvoe/gofakeit/v6"
)

// type aa struct {
// 	tt *int
// }

// type ab struct {
// 	abc []*aa
// }

func main() {
	Heap := FibonacciHeap{
		Count: setPoiter(0),
	}

	for i := 0; i < 20; i += 3 {
		Heap.Insert(setPoiter(HeapTree{
			Value: setPoiter(int(fake.Int32())),
			Order: setPoiter(0),
		}))
	}

	for i := 0; i < 20; i += 3 {
		v, err := Heap.ExtractMin()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(*v.Value)
		}
	}
}
