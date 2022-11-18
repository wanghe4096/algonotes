package main

/*
* 链表每k个节点逆序
* https://leetcode.cn/problems/reverse-nodes-in-k-group/description/
* 给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
* k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
* 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*
 */
import (
	"fmt"
)

type LNode struct {
	Next *LNode
	Val  int
}

func Append(head *LNode, e int) {
	p := head
	for p.Next != nil {
		p = p.Next
	}

	p.Next = &LNode{Next: nil, Val: e}

}

func PrintNode(p *LNode) {
	fmt.Printf("p=%p, p.Val=%d, p.Next=%p\n", p, p.Val, p.Next)
}

func PrintList(l *LNode) {
	p := l
	for p != nil {
		fmt.Printf("p=%p, p.Val=%d, p.Next=%p\n", p, p.Val, p.Next)
		p = p.Next
	}
}

func PrintE(l *LNode, k int) {
	p := l
	var i = 0
	for p != nil {
		if i%k == 0 {
			fmt.Print("| ")
		}
		fmt.Printf("%d ", p.Val)
		p = p.Next
		i++
	}
	fmt.Println()
}
func mockLinkList() *LNode {

	head := &LNode{Next: nil, Val: 0}

	for i := 1; i < 21; i++ {
		Append(head, i)
	}
	return head
}

func Len(head *LNode) int {
	p := head
	cnt := 0
	for p != nil {
		cnt++
		p = p.Next
	}
	return cnt
}

// 遍历链表， 每k个节点入栈， 用一个堆栈， 每个入栈， 然后出栈 , 空间复杂度O(n) , 时间复杂度O(1)
func ReverseK(head *LNode, k int) (res *LNode) {
	curr := head
	top := -1
	dumpy := &LNode{}
	p := dumpy
	i, N := 1, Len(head)
	stack := [5000]*LNode{}

	for curr != nil {
		if N%k != 0 && N-i+1 <= k {
			p.Next = curr
			break
		}

		top++
		stack[top] = curr
		curr = curr.Next
		//fmt.Printf("top=%d, p=%p, p.Val=%d, p.Next=%p\n", top, curr, curr.Val, curr.Next)
		if i%k == 0 {
			for top >= 0 {
				node := stack[top]
				p.Next = &LNode{Val: node.Val, Next: nil}
				p = p.Next
				top--
				PrintNode(p)
			}
			Pline()

		}
		i++
	}
	return dumpy.Next
}

func Pline() {
	fmt.Println("--------------------")
}
func main() {

	head := mockLinkList()
	//	PrintList(head)
	PrintE(head, 3)
	rl := ReverseK(head, 3)
	PrintE(head, 3)
	PrintList(rl)
	PrintE(rl, 3)
}
