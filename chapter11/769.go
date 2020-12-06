package main

import "fmt"

func main(){
	arr := []int{1,0,2,3,4}

	res := maxChunksToSorted(arr)

	fmt.Println(res)
}

func maxChunksToSorted(arr []int) int{
	if len(arr) < 2{
		return len(arr)
	}
	//目的是找出分割的数组中，最大值等于index的数，有的话就加1
	//count用来计算个数，cur用来表示最大的数

	count := 0
	cur := arr[0]
	for i,v := range arr {
		cur = max(cur,v)
		//如果最大的数与对应的下标相等，那么说明排序后是从小到大，能够链接起来的
		if cur == i{
			count++
		}
	}
	return count
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}