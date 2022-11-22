package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
 * 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
*
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	d1, d2 := l1, l2
	head := &ListNode{}
	result, p := head, d1
	dec := 0
	for d1 != nil && d2 != nil {
		s := d1.Val + d2.Val + dec
		dec = s / 10
		result.Next = &ListNode{Val: s % 10}
		result = result.Next
		d1 = d1.Next
		d2 = d2.Next
		p = p.Next

	}

	if d2 != nil {
		p = d2
	}

	for p != nil {
		s := (p.Val + dec)
		result.Next = &ListNode{Val: s % 10}
		dec = s / 10
		fmt.Printf("%d\t, p=%p, next=%p\n", p.Val, p, p.Next)
		p = p.Next
		result = result.Next
	}

	if dec > 0 {
		result.Next = &ListNode{Val: dec}
	}

	return head.Next
}

func initListNode(arr []int) *ListNode {
	head := &ListNode{}
	p := head
	for _, item := range arr {
		p.Next = &ListNode{Val: item, Next: nil}
		p = p.Next
	}

	return head.Next
}

func PrintList(l *ListNode) {
	p := l
	for p != nil {
		fmt.Printf("%d\t, p=%p, next=%p\n", p.Val, p, p.Next)
		p = p.Next
	}

	fmt.Println()
}
func main() {
	fmt.Println("vim-go")
	s1 := []int{1, 5, 3, 4}
	s2 := []int{4, 5, 6}
	l1 := initListNode(s1)
	l2 := initListNode(s2)
	res := addTwoNumbers(l1, l2)
	PrintList(res)

}
