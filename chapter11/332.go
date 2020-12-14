package main

import (
	"fmt"
	"sort"
)

func main(){
	tickets := [][]string{
		{"JFK","SFO"},
		{"JFK","ATL"},
		{"SFO","ATL"},
		{"ATL","JFK"},
		{"ATL","SFO"},
	}

	res := findItinerary(tickets)
	fmt.Println(res)
}

func findItinerary(tickets [][]string) []string {
	var (
		m  = map[string][]string{}
		res []string
	)

	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		m[src] = append(m[src], dst)
	}
	for key := range m {
		sort.Strings(m[key])
	}

	fmt.Println(m)

	var dfs func(curr string)
	dfs = func(curr string) {
		for {
			if v, ok := m[curr]; !ok || len(v) == 0 {
				break
			}
			tmp := m[curr][0]
			m[curr] = m[curr][1:]
			dfs(tmp)
		}
		//fmt.Println("curr",curr)
		//fmt.Println("res",res)
		res = append(res, curr)
	}

	dfs("JFK")

	//fmt.Println("res",res)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res) - 1 - i] = res[len(res) - 1 - i], res[i]
	}
	return res
}
