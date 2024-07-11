package heap

import "fmt"

func ExampleMaxHeap() {
	h := InitMaxHeap(10, 20, 30, 25, 5, 40, 35)
	for h.Size() > 0 {
		fmt.Println(h.Pop())
	}
	// Output:
	// 40
	// 35
	// 30
	// 25
	// 20
	// 10
	// 5
}
