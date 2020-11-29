package main

import "strconv"

func main(){

}

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}


type TreeNodeVal struct {
	Node *TreeNode
	Str string
}

func binaryTreePaths(root *TreeNode)  []string{
	ret := []string{}
	if nil == root{
		return ret
	}
	initRootNode := TreeNodeVal{
		Node:root,
		Str:strconv.Itoa(root.Val),
	}
	queue := []TreeNodeVal{
		initRootNode,
	}

	for len(queue) > 0{
		top := queue[0]
		queue = queue[1:]
		// 找到叶子节点
		if top.Node.Left == nil && top.Node.Right == nil{
			onePath := top.Str
			ret = append(ret,onePath)
		}
		// 入队左孩子
		if nil != top.Node.Left{
			queue = append(queue,TreeNodeVal{
				Node:top.Node.Left,
				Str :top.Str + "->" + strconv.Itoa(top.Node.Left.Val),
			})
		}
		// 入队右孩子
		if nil != top.Node.Right {
			queue = append(queue, TreeNodeVal{
				Node: top.Node.Right,
				Str: top.Str + "->" + strconv.Itoa(top.Node.Right.Val),
			})
		}
	}
	return ret
}