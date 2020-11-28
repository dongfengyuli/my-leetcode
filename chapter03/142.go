package main

import "fmt"

type ListNode struct {
	 Val int
	 Next *ListNode
}

var node4 = &ListNode{
	Val: -4,
	//Next:node2,
}
var node3 = &ListNode{
	Val: 0,
	Next:node4,
}
var node2 = &ListNode{
	Val: 2,
	Next:node3,
}
var node1 = &ListNode{
	Val: 3,
	Next:node2,
}

func main()  {

	res := detectCycle(node1)
	fmt.Println(res)

}


func detectCycle(head *ListNode) *ListNode{
	slow,fast := head,head
	for fast != nil{
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow{
			p := head
			for p != slow{
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

//a+(n+1)b+nc=2(a+b)⟹a=c+(n−1)(b+c)
//有了 a=c+(n-1)(b+c)a=c+(n−1)(b+c) 的等量关系，我们会发现：从相遇点到入环点的距离加上 n-1 圈的环长，恰好等于从链表头部到入环点的距离。
//因此，当发现 slow 与 fast 相遇时，我们再额外使用一个指针 ptr。起始，它指向链表头部；随后，它和 slow 每次向后移动一个位置。最终，它们会在入环点。

