package main

import "fmt"

func main(){
	haystack := "hello"
	needle := "ll"

	res := strStr(haystack,needle)
	fmt.Println(res)
}

func strStr(haystack string, needle string) int {
	if needle == "" {	//needle为空返回0
		return 0
	}
	if len(needle) > len(haystack) {	//needle长度大于haystack返回-1
		return -1
	}
	var i, j int	//定义指针i，j分别对应haystack和needle
	var index = -1 	//index位置初始值-1
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {	//如果当前指针处的字节相等
			if index == -1 {	//index是初始值-1，则将当前i记录下来
				index = i
			}
			if j == len(needle)-1 { //如果当前的j已经是needle的最后一位，则返回index
				return index
			}
			j++	//否则needle的指针j往后移动一位
		} else {	//当前指针处的字节不相等
			if index != -1 {
				i = index	//如果index不为初始值-1，则将i赋值为index，下次从index的下一个位置开始继续找
				index = -1
				j = 0		//再将index和j赋值回初始值
			}
		}
		i++	//不管上面怎么变动，haystack的指针每次都要往后移动一位
	}

	return -1
}

