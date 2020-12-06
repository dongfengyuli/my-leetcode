package main

import "fmt"

func main()  {
	obj := Constructor();
	obj.Push(3)
	obj.Push(4)
	obj.Push(9)
	obj.Push(21)

	fmt.Println("obj",obj)
	fmt.Println(obj.Pop())
	fmt.Println("obj",obj)
	obj.Push(24)

	fmt.Println("obj",obj)


}



type MyQueue struct {
	input []int
	output []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.input = append(this.input,x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.Peek()
	e := this.output[len(this.output) - 1]
	this.output = this.output[:len(this.output) - 1]
	return e
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.output) == 0 {
		for len(this.input) > 0 {
			this.output = append(this.output,this.input[len(this.input) - 1])
			this.input = this.input[:len(this.input) - 1]
		}
	}
	return this.output[len(this.output) - 1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.input) == 0 && len(this.output) == 0 {
		return true
	}
	return false
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */