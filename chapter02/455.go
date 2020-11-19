package main

import (
	"fmt"
	"sort"
)

func main()  {

	g := []int{1,2}
	s := []int{1,2,3}
	res := findContentChildren(g,s)
	fmt.Println(res)

}



func findContentChildren(g []int,s []int) int{
	sort.Ints(g)
	sort.Ints(s)
	i := 0
	for j := 0;i<len(g) && j < len(s);j++{
		if g[i] <= s[j]{
			i++
		}
	}
	return i
}