package main

import "fmt"

func main(){


	  //Your MyStack object will be instantiated and called as such:
		obj := Constructor4();
		obj.Push(5);
		obj.Push(2);
		obj.Push(15);
		obj.Push(8);
		obj.Push(35);
		obj.Push(21);
		obj.Push(35);
		obj.Push(23);
		fmt.Println("obj",obj)
		param_2 := obj.Pop();
		param_3 := obj.Top();
		param_4 := obj.IsEmpty();

		fmt.Println("param_2",param_2)
		fmt.Println("param_3",param_3)
		fmt.Println("param_4",param_4)

}


type MyStack struct {
	queue []int
}

func Constructor4() MyStack{
	return MyStack{}
}

func (this *MyStack) Push(x int){
	n := len(this.queue)
	this.queue = append(this.queue,x)
	for ;n>0;n--{
		this.queue = append(this.queue,this.queue[0])
		this.queue = this.queue[1:]
	}
}

func (this *MyStack) Pop()int{
	v := this.queue[0]
	this.queue = this.queue[1:]
	return v
}

func (this *MyStack) Top() int{
	return this.queue[0]
}

func (this *MyStack) IsEmpty() bool{
	return len(this.queue) == 0
}