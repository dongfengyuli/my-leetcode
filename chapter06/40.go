package main

import (
	"fmt"
	"sort"
)

func main()  {
	candidates := []int{10,1,2,7,6,1,5}
	target := 8
	res := combinationSum2(candidates,target)

	fmt.Println(res)
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates) //快排，懒得写
	res := [][]int{}
	dfs3(candidates, nil, target, 0, &res) //深度优先
	return res
}


func dfs3(candidates, nums []int, target, left int, res *[][]int){
	if target == 0{
		tmp := make([]int,len(nums))
		copy(tmp,nums)
		*res = append(*res,tmp)
		return
	}

	for i := left; i < len(candidates); i++ { // left限定，形成分支
		if i != left && candidates[i] == candidates[i-1] { // *同层节点 数值相同，剪枝
			continue
		}
		if target < candidates[i] { //剪枝
			return
		}

		dfs3(candidates, append(nums, candidates[i]), target-candidates[i], i+1, res) //*分支 i+1避免重复
	}
}
