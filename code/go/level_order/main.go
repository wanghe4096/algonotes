package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 * 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
 *
 */

type ListNode struct {
	Next *ListNode
	Prev *ListNode
	Data interface{}
}
type Queue struct {
	Head *ListNode // 虚拟头节点
	Tail *ListNode // 虚拟尾节点
	Size int
}

func InitQueue(queue *Queue) {
	queue.Size = 0
	queue.Head = &ListNode{}
	queue.Tail = &ListNode{}

	queue.Head.Next = queue.Tail
	queue.Head.Prev = nil

	queue.Tail.Next = nil
	queue.Tail.Prev = queue.Head

}

func (queue *Queue) Add(data interface{}) {
	head := queue.Head
	node := &ListNode{Data: data}
	next := head.Next
	head.Next = node
	node.Prev = head
	node.Next = next
	next.Prev = node
	queue.Size++
	return
}

// 从尾节点出队
func (queue *Queue) Pop() (res interface{}) {
	if queue.Size == 0 {
		return
	}

	res = queue.Tail.Prev.Data
	queue.Tail = queue.Tail.Prev
	queue.Tail.Next = nil
	queue.Size--
	return
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func get_nodes(queue *Queue) {

}
func levelOrder(root *TreeNode) [][]int {
	q := []*TreeNode{root}
	result := [][]int{}
	res := []int{}
	for len(q) > 0 {

		p := []*TreeNode{}
		for i := 0; i < len(q); i++ {
			res = append(res, q[i].Val)
			if q[i].Left != nil {
				p = append(p, q[i].Left)
			}

			if q[i].Right != nil {
				p = append(p, q[i].Right)
			}
		}

		result = append(result, res)
		q = p

	}
	return result

}
func main() {
	// first, test queue
	/*
		queue := &Queue{}
		InitQueue(queue)

		fmt.Println("add..")
		for i := 0; i < 10; i++ {
			queue.Add(i)
		}

		fmt.Println("pop...")
		for i := 0; i < 10; i++ {
			res := queue.Pop()
			fmt.Printf("i=%d\t res=%d \n", i, res)
		}
	*/

	// 测试层序遍历

	bt := &TreeNode{Val: 3}
	bt.Left = &TreeNode{Val: 9}
	bt.Right = &TreeNode{Val: 20}
	bt.Right.Left = &TreeNode{Val: 15}
	bt.Right.Right = &TreeNode{Val: 7}

	result := levelOrder(bt)

	for level, res := range result {
		fmt.Printf("level=%d, res=%+v\n", level, res)
	}

	PreOrder(bt)
	return
}

func PreOrder(bt *TreeNode) {
	if bt == nil {
		return
	}
	fmt.Println(bt.Val)
	PreOrder(bt.Left)
	PreOrder(bt.Right)
}
