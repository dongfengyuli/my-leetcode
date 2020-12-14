package main

import "fmt"

func main(){
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string)  int{
	if len(s) == 0{
		return 0
	}
	var freq = make(map[uint8]int,0)
	result := 0
	left,right := 0,-1
	for left < len(s){
		if right + 1 < len(s) && freq[s[right+1]] == 0 {
			freq[s[right+1]]++
			right++
		}else{
			freq[s[left]]--
			left++
		}
		result = max(result,right-left+1)
	}
	fmt.Println("freq",freq)
	return result
}


func max(a int,b int) int {
	if a >= b {
		return a
	}
	return b
}