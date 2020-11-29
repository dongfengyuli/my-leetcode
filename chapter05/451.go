package main

import (
	"fmt"
	"sort"
)

func main(){
	s := "tree"
	res := frequencySort(s)

	fmt.Println(res)
}


func frequencySort(s string) string {
	m := make(map[rune]int)
	ret := make([]rune,0)
	res := make([]rune,0)

	r := []rune(s)
	for _,v := range r {
		i,ok := m[v]
		if ok {
			m[v] = i + 1
		}else{
			m[v] = 1
			ret = append(ret,v)
		}
	}

	fmt.Println("m",m)
	fmt.Println("ret",ret)

	sort.Slice(ret,func(i,j int) bool{
		return m[ret[i]] > m[ret[j]]
	})

	fmt.Println("mm",m)
	fmt.Println("retret",ret)


	for i:=0;i<len(ret);i++{
		for j:=0;j<m[ret[i]];j++{
			res = append(res,ret[i])
		}
	}
	return string(res)
}