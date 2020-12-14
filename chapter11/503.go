package main

import "fmt"

func main(){
	nums := []int{3,8,4,1,2}
	res := nextGreaterElements(nums)

	fmt.Println("res",res)
}



//单调栈
func nextGreaterElements(nums []int) []int {
	if len(nums) < 1  {
		return []int{}
	}
	// 拼上一个相同的数组，便于统一处理
	nums=append(nums,nums...)
	stack:=[]int{0}

	for i :=1;i<len(nums);i++ {
		//fmt.Println(nums,i)
		//比栈里的小 插入栈
		if nums[i] <= nums[stack[len(stack)-1]] {
			stack = append(stack,i)
			continue
		}
		// 比栈里的 大 可以出栈
		for len(stack) != 0 && nums[i] >nums[stack[len(stack)-1]] {
			nums[stack[len(stack)-1]] = nums[i]
			stack = stack[:len(stack)-1]
		}
		// 提前结束的条件
		if len(stack) == 0 && i>=len(nums)/2{
			break
		}
		// 把当前的插入栈
		stack = append(stack,i)
	}

	//fmt.Println("stack",stack)
	//fmt.Println("nums",nums)

	// 栈里还有的元素 也就是 没有比它大的 置为-1吧
	for e:=range stack{
		// 提前结束的条件
		if stack[e]>=len(nums)/2{
			break
		}
		nums[stack[e]]=-1
	}
	return nums[0:len(nums)/2]
}

