package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/intersection-of-two-linked-lists/
// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
//

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB

	lenOfListFunc := func(head *ListNode) int {
		p := head
		cnt := 0
		for p != nil {
			cnt++
			p = p.Next
		}
		return cnt
	}

	p1_len, p2_len := lenOfListFunc(p1), lenOfListFunc(p2)
	skip := p1_len - p2_len
	if skip < 0 {
		p1, p2 = p2, p1
		skip = 0 - skip
	}

	for i := 0; i < skip; i++ {
		p1 = p1.Next

	}

	for p1 != nil && p2 != nil {
		if p1 == p2 {
			return p1
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return nil
}

func main() {
	head := &ListNode{Val: 0, Next: nil}

	p := head

	for i := 1; i < 5; i++ {
		p.Next = &ListNode{Val: i}
		p = p.Next
	}

	head2 := &ListNode{}

	p2 := head2
	for i := 10; i < 20; i++ {
		p2.Next = &ListNode{Val: i}
		p2 = p2.Next
	}

	p2.Next = p
	p2 = p2.Next

	for i := 30; i < 40; i++ {
		p.Next = &ListNode{Val: i}
		p = p.Next
	}

	p = head
	i := 0
	for p != nil {
		i++
		fmt.Printf("%d\tval=%d, p = %p, next = %p \n", i, p.Val, p, p.Next)
		p = p.Next
	}

	p2 = head2
	i = 0
	for p2 != nil {
		i++
		fmt.Printf("%d\tval2=%d, p2 = %p, next = %p \n", i, p2.Val, p2, p2.Next)
		p2 = p2.Next
	}

	p, p2 = head, head2
	node := getIntersectionNode(p, p2)
	if node != nil {
		fmt.Printf("p=%p, next=%p, val=%d\n", node, node.Next, node.Val)
	} else {
		fmt.Println("node=nil")
	}

}
