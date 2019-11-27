// not a good implementation heap data structure, just for practise
package main

// importing fmt package and container/heap
import (
	"fmt"
)

// integerHeap a type
type IntegerHeap []int

// IntegerHeap method - gets the length of integerHeap
func (iheap IntegerHeap) Len() int { return len(iheap) }

// IntegerHeap method - checks if element of i index is less than j index
func (iheap IntegerHeap) Less(i, j int) bool { return iheap[i] < iheap[j] }

// IntegerHeap method -swaps the element of i to j index
func (iheap IntegerHeap) Swap(i, j int) { iheap[i], iheap[j] = iheap[j], iheap[i] }

//IntegerHeap method -pushes the item
func (iheap *IntegerHeap) Push(heapintf interface{}) {

	*iheap = append(*iheap, heapintf.(int))
	isSwap := false
	lastIdx := iheap.Len() - 1
	for lastIdx > -1 {
		pIdx := iheap.findParentIdx(lastIdx)
		if pIdx > -1 {
			if iheap.Less(lastIdx, pIdx) {
				iheap.Swap(lastIdx, pIdx)
				isSwap = true
			}
		}
		if isSwap {
			lastIdx = pIdx
		} else {
			lastIdx = -1
		}
	}

}

func (iheap *IntegerHeap) findParentIdx(idx int) int {
	if idx == 0 {
		return -1
	}

	if idx%2 == 0 {
		return (idx - 2) / 2
	}

	return (idx - 1) / 2
}

// initialize the heap
func (iheap IntegerHeap) Init() {
	n := iheap.Len()
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if iheap.Less(j, i) {
				iheap.Swap(j, i)
			}
		}
	}
}

func (iheap IntegerHeap) HeapfyDown(idx int) int {
	if idx*2+2 < iheap.Len() {
		if iheap.Less(idx*2+1, idx*2+2) {
			if iheap.Less(idx*2+1, idx) {
				iheap.Swap(idx*2+1, idx)
				return idx*2 + 1
			}

			return -1
		} else if iheap.Less(idx*2+2, idx*2+1) {
			if iheap.Less(idx*2+2, idx) {
				iheap.Swap(idx*2+2, idx)
				return idx*2 + 2
			}
			return -1
		}
	} else if (idx*2+1) < iheap.Len() && (idx*2+2) >= iheap.Len() {
		if iheap.Less(idx*2+1, idx) {
			iheap.Swap(idx*2+1, idx)
		}

		return -1
	}
	return -1
}

// pop out the root element
func (iheap *IntegerHeap) PopMin() int {
	n := iheap.Len()
	iheap.Swap(0, n-1)
	ret := (*iheap)[n-1]
	*iheap = (*iheap)[:n-1]
	idx := 0
	for idx > -1 {
		idx = iheap.HeapfyDown(idx)
	}
	return ret
}

//IntegerHeap method -pops the item from the heap
func (iheap *IntegerHeap) Pop() interface{} {
	var n int
	var x1 int
	var previous IntegerHeap = *iheap
	n = len(previous)
	x1 = previous[n-1]
	*iheap = previous[0 : n-1]
	return x1
}

func (iheap IntegerHeap) ShowAll() {
	for i, v := range iheap {
		fmt.Println("===", i, v)
	}
}

// main method
func main() {
	var intHeap *IntegerHeap = &IntegerHeap{12, 7, 17, 3, 10, 13, 1}

	// container/heap
	// heap.Init(intHeap)
	// heap.Push(intHeap, 2)
	// fmt.Printf("minimum: %d\n", (*intHeap)[0])

	// self implement heap
	intHeap.Init()
	intHeap.ShowAll()

	fmt.Println("==============")
	intHeap.Push(2)
	intHeap.ShowAll()
	fmt.Println("===============")
	fmt.Printf("minium: %d\n", intHeap.PopMin())
	fmt.Println("===============")
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", intHeap.Pop())
	}
}
