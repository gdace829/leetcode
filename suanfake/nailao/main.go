package main

import (
	"container/heap"
	"fmt"
)

type nailao []uint32

func (n nailao) Less(i, j int) bool {
	a := n[i]
	cntA := 0
	b := n[j]
	cntB := 0
	for a != 0 {
		if a&1 == 1 {
			cntA++
		}
		a >>= 1
	}
	for b != 0 {
		if b&1 == 1 {
			cntB++
		}
		b >>= 1
	}
	if cntA > cntB {
		return true
	} else if cntA < cntB {
		return false
	} else {
		return n[i] > n[j]
	}
}
func (pq nailao) Len() int {
	return len(pq)
}

func (pq nailao) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *nailao) Push(x interface{}) {
	*pq = append(*pq, x.(uint32))
}

func (pq *nailao) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func main() {
	var n int
	var option string
	nailao := &nailao{}
	heap.Init(nailao)
	fmt.Scan(&n)
	for n > 0 {
		fmt.Scan(&option)
		if option == "A" {
			var a uint32
			fmt.Scan(&a)
			heap.Push(nailao, a)
		} else if option == "B" {
			heap.Pop(nailao)
		} else {
			fmt.Println((*nailao)[0])
		}
		n--
	}

}
