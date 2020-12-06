package main

import (
	"fmt"
	"math"
)

func main(){
	obj := Constructor2()
	obj.Push(2)
	obj.Push(4)
	obj.Push(21)
	obj.Push(15)
	obj.Push(13)
	obj.Push(27)
	obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.GetMin()

	fmt.Println("obj",obj)
	fmt.Println("param_3",param_3)
	fmt.Println("param_4",param_4)
}



type MinStack struct {
	stack []int
	minStack []int
}

func Constructor2() MinStack {
	return MinStack{
		stack:[]int{},
		minStack:[]int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int){
	this.stack = append(this.stack,x)
	top := this.minStack[len(this.minStack) - 1]
	this.minStack = append(this.minStack,min(x,top))
}

func (this *MinStack) Pop(){
	this.stack = this.stack[:len(this.stack) -1]
	this.minStack = this.minStack[:len(this.minStack) -1]
}

func (this *MinStack) Top() int{
	return this.stack[len(this.stack) -1]
}

func (this *MinStack) GetMin() int{
	return this.minStack[len(this.minStack) -1]
}





func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}