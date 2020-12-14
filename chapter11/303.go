package main

import "fmt"

func main(){
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor3(nums);
	param_1 := obj.SumRange(2,5);

	fmt.Println("obj",obj)
	fmt.Println("param_1",param_1)
}



type NumArray struct {
	data []int
}

func Constructor3(nums []int) NumArray{
	for i:=1;i<len(nums);i++{
		nums[i] += nums[i-1]
	}
	return NumArray{data:nums}
}


func (this *NumArray) SumRange(i int,j int) int{
	if 0 == i{
		return this.data[j]
	}
	return this.data[j] - this.data[i-1]
}