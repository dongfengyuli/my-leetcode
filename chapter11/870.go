package main

import (
	"fmt"
	"sort"
)

func main(){
	A := []int{12,24,8,32}
	B := []int{13,25,32,11}

	res := advantageCount(A,B)
	fmt.Println(res)
}


type node struct{
	num int
	k int
}
type nodes []node

func (n nodes)Len()int{
	return len(n)
}

func (n nodes)Less(i,j int)bool{
	return n[i].num>n[j].num
}

func (n nodes)Swap(i,j int){
	n[i],n[j]=n[j],n[i]
}

func advantageCount(A []int, B []int) []int {

	var a=make([]node,len(A) )
	var b=make([]node,len(A) )
	for i:=0;i<len(A);i++{
		a[i]=node{A[i],i}
		b[i]=node{B[i],i}
	}
	sort.Sort(nodes(a))
	sort.Sort(nodes(b))
	fmt.Println("a:",a)
	fmt.Println("b:",b)
	var x int =len(a)-1
	var res =make([]int,len(A))
	var j,index,i int
	for ; index < len(A) ; index++ {
		if a[i].num>b[j].num{
			res[b[j].k]=a[i].num
			i++
		}else {
			res[b[j].k]=a[x].num
			x--
		}
		j++
	}
	return res
}
