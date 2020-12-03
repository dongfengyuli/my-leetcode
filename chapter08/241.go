package main

import (
	"fmt"
	"strconv"
)

func main(){
	input :=  "2*3-4*5"
	res := diffWaysToCompute(input)
	fmt.Println(res)
}

func diffWaysToCompute(input string) []int{
	// 如果是数字，直接返回
	if isDigit(input){
		tmp,_ := strconv.Atoi(input)
		return []int{tmp}
	}
	// 空切片
	var res []int
	// 遍历字符串
	for index,c := range input{
		tmpC := string(c)
		if tmpC == "+" || tmpC == "-" || tmpC == "*"{
			// 如果是运算符，则计算左右两边的算式
			left := diffWaysToCompute(input[:index])
			right := diffWaysToCompute(input[index+1:])
			for _,leftNum := range left{
				for _,rightNum := range right{
					var addNum int
					if tmpC == "+"{
						addNum = leftNum + rightNum
					}else if tmpC == "-"{
						addNum = leftNum - rightNum
					}else{
						addNum = leftNum * rightNum
					}
					res = append(res,addNum)
				}
			}
			//
			//
		}

	}

	return res
}


// 判断是否为全数字
func isDigit(input string) bool {
	_, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	return true
}