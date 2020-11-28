package main

import "fmt"

func main(){
	s := "abpcplea"
	d := []string{"ale","apple","monkey","plea"}
	res := findLongestWord(s,d)
	fmt.Println(res)
}

func findLongestWord(s string,d []string)  string{
	res := ""
	if len(s) <= 0{
		return res
	}
	for i := range d{
		if subString(s,d[i]){
			if len(d[i]) > len(res) || (len(d[i]) == len(res) && d[i] < res ){
				res = d[i]
			}
		}
	}
	return res
}

func subString(s string,sub string)  bool{
	if len(sub) <= 0{
		return true
	}
	p1 := 0
	p2 := 0
	for p1 < len(s){
		if s[p1] == sub[p2]{
			p1++
			p2++
		}else{
			p1++
		}
		if p2 == len(sub){
			return true
		}
	}
	return false
}