package main

import "fmt"

func main(){
	s := "anagram"
	t := "nagaram"
	res := isAnagram(s,t)
	fmt.Println(res)
}

func isAnagram(s string,t string) bool{
	if len(s) != len(t){
		return false
	}

	strMap := make(map[string]int)

	for i:=0;i<len(s);i++{
		strMap[string(s[i])]++
	}

	for j:=0;j<len(t);j++{
		strMap[string(s[j])]--
	}

	for _,v := range strMap{
		if v != 0{
			return false
		}
	}
	return true
}


