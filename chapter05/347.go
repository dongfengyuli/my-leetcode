package main

import (
	"container/heap"
	"fmt"
)

func main(){
	nums := []int{1,1,1,2,2,3}
	k := 2

	res := topKFrequent(nums,k)
	fmt.Println(res)
}

type Feq struct{
	val int
	cnt int
}

type FeqMinHeap []Feq

func (pq *FeqMinHeap) Len() int{
	return len(*pq)
}

func (pq *FeqMinHeap) Less(i,j int) bool{
	return (*pq)[i].cnt < (*pq)[j].cnt
}

func (pq *FeqMinHeap) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *FeqMinHeap) Push(x interface{}) {
	*pq = append(*pq, x.(Feq))
}

func (pq *FeqMinHeap) Pop() interface{} {
	n := len(*pq) - 1
	x := (*pq)[n]
	*pq = (*pq)[:n]
	return x
}
func (pq *FeqMinHeap) Peek() Feq {
	return (*pq)[0]
}

func topKFrequent(nums []int,k int) []int{
	cntMap := make(map[int]int)
	for _,num := range nums{
		cntMap[num]++
	}
	minHeap := &FeqMinHeap{}
	for key,val := range cntMap{
		heap.Push(minHeap,Feq{
			val:key,
			cnt:val,
		})
		if minHeap.Len() > k{
			heap.Pop(minHeap)
		}
	}
	res := make([]int,k)
	i := 0
	for minHeap.Len() > 0{
		res[i] = minHeap.Pop().(Feq).val
		i++
	}
	return res
}