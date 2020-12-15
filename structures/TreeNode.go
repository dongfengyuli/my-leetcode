package structures

import (
	"fmt"
)

// TreeNode is tree's node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NULL 方便添加测试数据
var NULL = -1 << 63

// Ints2TreeNode 利用 []int 生成 *TreeNode
func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// GetTargetNode 返回 Val = target 的 TreeNode
// root 中一定有 node.Val = target
func GetTargetNode(root *TreeNode, target int) *TreeNode {
	if root == nil || root.Val == target {
		return root
	}

	res := GetTargetNode(root.Left, target)
	if res != nil {
		return res
	}
	return GetTargetNode(root.Right, target)
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	msg := fmt.Sprintf("%d 不存在于 %v 中", val, nums)
	panic(msg)
}

// PreIn2Tree 把 preorder 和 inorder 切片转换成 二叉树
func PreIn2Tree(pre, in []int) *TreeNode {
	if len(pre) != len(in) {
		panic("preIn2Tree 中两个切片的长度不相等")
	}

	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: pre[0],
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)

	res.Left = PreIn2Tree(pre[1:idx+1], in[:idx])
	res.Right = PreIn2Tree(pre[idx+1:], in[idx+1:])

	return res
}

// InPost2Tree 把 inorder 和 postorder 切片转换成 二叉树
func InPost2Tree(in, post []int) *TreeNode {
	if len(post) != len(in) {
		panic("inPost2Tree 中两个切片的长度不相等")
	}

	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: post[len(post)-1],
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)

	res.Left = InPost2Tree(in[:idx], post[:idx])
	res.Right = InPost2Tree(in[idx+1:], post[idx:len(post)-1])

	return res
}

// Tree2Preorder 把 二叉树 转换成 preorder 的切片
func Tree2Preorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := []int{root.Val}
	res = append(res, Tree2Preorder(root.Left)...)
	res = append(res, Tree2Preorder(root.Right)...)

	return res
}

// Tree2Inorder 把 二叉树转换成 inorder 的切片
func Tree2Inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := Tree2Inorder(root.Left)
	res = append(res, root.Val)
	res = append(res, Tree2Inorder(root.Right)...)

	return res
}

// Tree2Postorder 把 二叉树 转换成 postorder 的切片
func Tree2Postorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := Tree2Postorder(root.Left)
	res = append(res, Tree2Postorder(root.Right)...)
	res = append(res, root.Val)

	return res
}

// Equal return ture if tn == a
func (tn *TreeNode) Equal(a *TreeNode) bool {
	if tn == nil && a == nil {
		return true
	}

	if tn == nil || a == nil || tn.Val != a.Val {
		return false
	}

	return tn.Left.Equal(a.Left) && tn.Right.Equal(a.Right)
}

// Tree2ints 把 *TreeNode 按照行还原成 []int
func Tree2ints(tn *TreeNode) []int {
	res := make([]int, 0, 1024)

	queue := []*TreeNode{tn}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			nd := queue[i]
			if nd == nil {
				res = append(res, NULL)
			} else {
				res = append(res, nd.Val)
				queue = append(queue, nd.Left, nd.Right)
			}
		}
		queue = queue[size:]
	}

	i := len(res)
	for i > 0 && res[i-1] == NULL {
		i--
	}

	return res[:i]
}

// T2s convert *TreeNode to []int
func T2s(head *TreeNode, array *[]int) {
	fmt.Printf("运行到这里了 head = %v array = %v\n", head, array)
	// fmt.Printf("****array = %v\n", array)
	// fmt.Printf("****head = %v\n", head.Val)
	*array = append(*array, head.Val)
	if head.Left != nil {
		T2s(head.Left, array)
	}
	if head.Right != nil {
		T2s(head.Right, array)
	}
}

//104. 二叉树的最大深度

func MaxDepth(root *TreeNode) int{
	if nil == root{
		return 0
	}

	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)

	if left >= right{
		return left + 1
	}else {
		return right + 1
	}
}

//110. 平衡二叉树

func IsBalanced(root *TreeNode) bool{
	if nil == root{
		return true
	}
	//fmt.Println("MaxDepth(root.Left)",MaxDepth(root.Left))
	//fmt.Println("MaxDepth(root.Right)",MaxDepth(root.Right))
	diff := MaxDepth(root.Left) - MaxDepth(root.Right)
	if diff > 1 || diff < -1{
		return false
	}
	return IsBalanced(root.Left) && IsBalanced(root.Right)
}

//543. 二叉树的直径

func DiameterOfBinaryTree(root *TreeNode) int{
	ans := 0
	if nil == root{
		return ans
	}

	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)
	if left + right > ans{
		ans = left + right
	}
	return ans
}

//437. 路径总和 III

func PathSum(root *TreeNode,sum int) int{
	if nil == root{
		return 0
	}
	return pathSumFrom(root,sum) + PathSum(root.Left,sum) + PathSum(root.Right,sum)
}

func pathSumFrom(node *TreeNode,sum int) int{
	if nil == node{
		return 0
	}
	ans := 0
	if node.Val == sum {
		ans = 1
	}
	return ans + pathSumFrom(node.Left,sum - node.Val) + pathSumFrom(node.Right,sum-node.Val)
}

//101. 对称二叉树

func IsSymmetric(root *TreeNode) bool{
	return nil == root || isSame(root.Left,root.Right)
}

func isSame(left,right *TreeNode) bool{
	if nil == left || nil ==right{
		return left == right
	}else if left.Val != right.Val{
		return false
	}else{
		return isSame(left.Left,right.Right) && isSame(left.Right,right.Left)
	}
}

//二叉树的前序遍历
//方法一：递归

/*func PreOrderTraversal(root *TreeNode) []int{
	res := []int{}
	if root == nil{
		return res
	}
	res = append(res,root.Val)
	res = append(res,PreOrderTraversal(root.Left)...)
	res = append(res,PreOrderTraversal(root.Right)...)
	return res
}*/

//方法二：迭代
func PreOrderTraversal(root *TreeNode) []int{
	if root == nil{
		return nil
	}

	var stack []*TreeNode
	stack = append(stack,root)

	var res []int
	for len(stack) > 0{
		p := stack[len(stack) -1]
		stack = stack[0:len(stack) -1]
		res = append(res,p.Val)
		if p.Right != nil{
			stack = append(stack,p.Right)
		}
		if p.Left != nil {
			stack = append(stack,p.Left)
		}
	}
	return res
}

//94. 二叉树的中序遍历
//方法一：递归

/*var res []int
func InOrderTraversal(root *TreeNode) []int {
	//res = make([]int, 0)
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
	}
}*/


//方法二：迭代

func InOrderTraversal(root *TreeNode) []int{
	if root == nil{
		return nil
	}

	var stack []*TreeNode
	var res []int

	for root != nil || len(stack) > 0{
		for root != nil{
			stack = append(stack,root)
			root = root.Left
		}
		root = stack[len(stack) -1]
		stack = stack[:len(stack) -1]
		res = append(res,root.Val)
		root = root.Right
	}
	return res
}

//145. 二叉树的后序遍历

/*var res []int
func PostOrderTraversal(root *TreeNode) []int {
	//res = make([]int, 0)
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		dfs(root.Right)
		res = append(res, root.Val)
	}
}*/


func PostOrderTraversal(root *TreeNode) []int{
	if root == nil{
		return nil
	}
	res := []int{}
	stack := []*TreeNode{root}
	var node *TreeNode

	for len(stack) > 0{
		node,stack = stack[len(stack) - 1],stack[:len(stack) -1]
		//fmt.Println("res","node","node.Val",res,node,node.Val)
		res = append([]int{node.Val},res...)
		if node.Left != nil{
			stack = append(stack,node.Left)
		}
		if node.Right != nil {
			stack = append(stack,node.Right)
		}
	}
	return res
}

//102. 二叉树的层序遍历