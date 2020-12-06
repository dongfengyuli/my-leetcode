package main

import (
	"fmt"
	"strings"
)

func main(){
	words := []string{"abcw","baz","foo","bar","xtfn","abcdef"}

	res := maxProduct(words)
	fmt.Println(res)
}

/*func maxProduct(words []string) int {
	if words == nil || len(words) < 2 {
		return 0
	}
	length := len(words)
	masks := make([]int, length)
	for i, word := range words[:] {
		for _, c := range word {
			masks[i] |= 1 << uint(c - 'a')
		}
	}

	fmt.Println("masks",masks)

	max := 0
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if (masks[i] & masks[j]) == 0 {
				prod := len(words[i]) * len(words[j])
				if prod > max {
					max = prod
				}
			}
		}
	}
	return max
}
*/


func maxProduct(words []string) int {
	lenth := len(words)
	var res = 0
	for i:=0;i<lenth-1;i++{
		for j:=i+1;j<lenth;j++{
			if !strings.ContainsAny(words[i],words[j]) && res < len(words[i]) * len(words[j]){
				res = len(words[i]) * len(words[j])
			}
		}
	}
	return res
}


