package main

import "fmt"

func main(){
	s := "abccccdd"
	res := longestPalindrome(s)
	fmt.Println(res)
}


func longestPalindrome(s string) int{
	dict := make(map[byte]int,0)
	for i:=0;i<len(s);i++{
		dict[s[i]]++
	}

	odd := 0
	for _,v := range dict{
		if v%2 != 0{
			odd++
		}
	}

	ans := len(s)
	if 0 != odd{
		ans -= odd - 1
	}
	return ans
}
