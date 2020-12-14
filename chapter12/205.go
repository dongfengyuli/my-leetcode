package main

import "fmt"

func main(){
	s := "egg"
	t := "add"

	res := isIsomorphic(s,t)
	fmt.Println(res)
}



func isIsomorphic(s string,t string) bool{
	length := len(s)
	num1 := make([]int,256)
	num2 := make([]int,256)

	for i:=0;i<length;i++{
		if num1[s[i]] != num2[t[i]]{
			return false
		}
		num1[s[i]] = i + 1
		num2[t[i]] = i + 1
	}
	//fmt.Println("num1",num1)
	//fmt.Println("num2",num2)
	return true
}