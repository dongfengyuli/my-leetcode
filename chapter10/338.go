package main

import "fmt"

func main(){
	res := countBits(12)
	fmt.Println(res)
}


/*func countBits(num int) []int {
	arr:=make([]int,num+1)
	for i:=1;i<=num;i++{
		arr[i]=arr[i&(i-1)]+1
	}
	return arr
}
*/


func countBits(num int) []int {
	d:=make([]int,num+1)
	d[0]=0
	for i:=1;i<=num;i++{
		if i&1==1{
			d[i]=d[i-1]+1
		}else{
			d[i]=d[i/2]
		}
	}
	return d
}
