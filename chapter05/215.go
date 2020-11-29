package main

import (
	"container/heap"
	"fmt"
)

func main(){
	nums := []int{3,2,1,5,6,4}
	k := 2
	res := findKthLargest(nums,k)
	fmt.Println(res)
}



func findKthLargest(nums []int, k int) int {
	m := make(TopList,0)
	size := 0
	for i:= range nums{
		if size < k{
			heap.Push(&m,nums[i])
			size++
		}else{
			if nums[i] > m[0]{
				heap.Pop(&m)
				heap.Push(&m,nums[i])
			}
		}
	}
	/*fmt.Println("####")
	fmt.Println(m)
	fmt.Println("####")*/
	return m[0]
}


type TopList []int

func (t TopList)Len() int{
	return len(t)
}

func (t TopList)Swap(i,j int){
	t[i],t[j] = t[j],t[i]
}

func (t TopList)Less(i,j int) bool {
	return t[i] < t[j]
}

func (t *TopList)Push(x interface{}){
	*t = append(*t,x.(int))
}



func (t *TopList)Pop() interface{}{
	x := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return x
}