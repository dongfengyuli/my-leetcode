package structures

import (
	"fmt"
)

// ListNode 是链接节点
// 这个不能复制到*_test.go文件中。会导致Travis失败
type ListNode struct {
	Val  int
	Next *ListNode
}

// List2Ints convert List to []int
func List2Ints(head *ListNode) []int {
	// 链条深度限制，链条深度超出此限制，会 panic
	limit := 100

	times := 0

	res := []int{}
	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("链条深度超过%d，可能出现环状链条。请检查错误，或者放宽 l2s 函数中 limit 的限制。", limit)
			panic(msg)
		}

		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// Ints2List convert []int to List
func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

// GetNodeWith returns the first node with val
func (l *ListNode) GetNodeWith(val int) *ListNode {
	res := l
	for res != nil {
		if res.Val == val {
			break
		}
		res = res.Next
	}
	return res
}

// Ints2ListWithCycle returns a list whose tail point to pos-indexed node
// head's index is 0
// if pos = -1, no cycle
func Ints2ListWithCycle(nums []int, pos int) *ListNode {
	head := Ints2List(nums)
	if pos == -1 {
		return head
	}
	c := head
	for pos > 0 {
		c = c.Next
		pos--
	}
	tail := c
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = c
	return head
}

//206 反转链表

func ReverseList(head *ListNode) *ListNode{
	var newHead *ListNode
	for head != nil{
		nextNode := head.Next
		head.Next = newHead
		newHead = head
		head = nextNode
	}
	return newHead
}

//21合并有序链表

func MergeTwoLists(l1 *ListNode,l2 *ListNode) *ListNode{
	if nil == l1{
		return l2
	}
	if nil == l2{
		return l1
	}

	if l1.Val < l2.Val{
		l1.Next = MergeTwoLists(l1.Next,l2)
		return l1
	}else{
		l2.Next = MergeTwoLists(l1,l2.Next)
		return l2
	}
}

//24. 两两交换链表中的节点
func SwapPairs(head *ListNode) *ListNode{
	dummyHead := &ListNode{0,head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil{
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}

//160. 相交链表

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if nil == headA || nil == headB{
		return nil
	}

	ptrA,ptrB := headA,headB
	for ptrA != ptrB{
		if nil == ptrA{
			ptrA = headB
		}else{
			ptrA = ptrA.Next
		}

		if nil == ptrB{
			ptrB = headA
		}else{
			ptrB = ptrB.Next
		}
	}
	return ptrA
}

//234. 回文链表
func Ispalindrome(head *ListNode) bool{
	nums := make([]int,0,128)
	for nil != head{
		nums = append(nums,head.Val)
		head = head.Next
	}

	for i,j := 0,len(nums) - 1;i<j;{
		if nums[i] != nums[j]{
			return false
		}
		i++
		j--
	}
	return true
}

//83. 删除排序链表中的重复元素

func DeleteDuplicates(head *ListNode) *ListNode{
	if nil == head || nil == head.Next{
		return head
	}

	prePtr := head
	currentPtr := head.Next
	lastValue := prePtr.Val

	for nil != currentPtr{
		if currentPtr.Val != lastValue{
			lastValue = currentPtr.Val
			prePtr = currentPtr
			currentPtr = currentPtr.Next
		}else{
			prePtr.Next = currentPtr.Next
			currentPtr = currentPtr.Next
		}
	}
	return head
}

//328. 奇偶链表
func OddEvenList(head *ListNode) *ListNode{
	if head == nil {
		return head
	}
	odd := head
	even := head.Next
	evenHead := even

	for even != nil && even.Next != nil{
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

//19. 删除链表的倒数第N个节点

func RemoveNthFromEnd(head *ListNode,n int) *ListNode{
	result := &ListNode{}
	result.Next = head
	var pre *ListNode
	cur := result
	fmt.Println("cur:","result:",cur,result)
	i := 1
	for head != nil{
		if i >= n{
			pre = cur
			cur = cur.Next
		}
		head = head.Next
		i++
	}
	pre.Next = pre.Next.Next
	return result.Next
}

//148. 排序链表

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

//141. 环形链表
func HasCycle(head *ListNode) bool{
	if head == nil {
		return false
	}
	fast,slow := head,head

	for fast != nil && fast.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow{
			return true
		}
	}
	return false
}

//142. 环形链表 II
func DelectCycle(head *ListNode) *ListNode{
	slow,fast := head,head
	for fast != nil{
		slow = slow.Next
		if fast.Next == nil{
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