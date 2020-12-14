package main

import "fmt"

func main(){
	nums := []int{1, 3, 5,7,11}
	obj := Constructor(nums);

	fmt.Println("obj",obj)

	obj.Update(2,15);
	param_2 := obj.SumRange(1,4);

	fmt.Println(param_2)
}

type NumArray struct {
	Nums  []int
	Tree  []int
}
func Constructor(nums []int) NumArray {
	this := NumArray{
		Nums:make([]int,len(nums)),
		Tree:make([]int,len(nums)+1),
	}
	for i := 0; i < len(nums);i++{
		this.Update(i,nums[i])
	}
	return this
}
func (this *NumArray) Update(i int, val int)  {
	diff := val - this.Nums[i]
	this.Nums[i] = val
	for j := i+1; j <= len(this.Nums); j += j & (-j){
		this.Tree[j] += diff
	}
}
func (this *NumArray) SumRange(i int, j int) int {
	return this.sum(j+1) - this.sum(i)
}
func (this *NumArray) sum(k int)int{
	sum := 0
	for i := k; i > 0;i -= i &(-i){
		sum += this.Tree[i]
	}
	return sum
}




//解题思路:
//1,初始化的时候 num数组一定是开辟一个新空间的
//2,i += i & (-i) 与 j -= j & (-j) 是树状数组的下标移动方法 与二叉树 2 * idx + 1 和 2 * idx + 2 是一样的
//3,尽量想象 数状数组构成的图形



