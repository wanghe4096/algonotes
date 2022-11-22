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

type ListNode struct {
	Next *ListNode
	Val  int
}

// 往链表里追加节点
func Append(head *ListNode, e int) {
	p := head
	for p.Next != nil {
		p = p.Next
	}

	p.Next = &ListNode{Next: nil, Val: e}

}

func PrintNode(p *ListNode) {
	fmt.Printf("p=%p, p.Val=%d, p.Next=%p\n", p, p.Val, p.Next)
}

func PrintList(l *ListNode) {
	p := l
	for p != nil {
		fmt.Printf("p=%p, p.Val=%d, p.Next=%p\n", p, p.Val, p.Next)
		p = p.Next
	}
}

func PrintE(l *ListNode, k int) {
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

func mockLinkList(arr ...int) *ListNode {

	head := &ListNode{Next: nil, Val: 0}

	for _, item := range arr {
		Append(head, item)
	}
	return head.Next
}

func Len(head *ListNode) int {
	p := head
	cnt := 0
	for p != nil {
		cnt++
		p = p.Next
	}
	return cnt
}

// 遍历链表， 每k个节点入栈， 用一个堆栈， 每个入栈， 然后出栈 , 空间复杂度O(n) , 时间复杂度O(1)
func ReverseK(head *ListNode, k int) (res *ListNode) {
	curr := head
	top := -1
	dumpy := &ListNode{}
	p := dumpy
	i, N := 1, Len(head)
	stack := [5000]*ListNode{}

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
				p.Next = &ListNode{Val: node.Val, Next: nil}
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

func LenOfList(head *ListNode) int {
	p := head
	cnt := 0
	for p != nil {
		cnt++
		p = p.Next
	}

	return cnt
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = myReverse(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}

/*
作者：力扣官方题解
链接：https://leetcode.cn/problems/reverse-nodes-in-k-group/solutions/248591/k-ge-yi-zu-fan-zhuan-lian-biao-by-leetcode-solutio/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

/*
func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	prev := hair
	for head != nil {
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		next := tail.Next
		head, tail := myReverse(head, tail)
		prev.Next = head
		tail.Next = next
		prev = tail
		head = tail.Next
	}

	return hair.Next
}

func myReverse(head, tail *ListNode) (h, t *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}*/

func Pline() {
	fmt.Println("--------------------")
}

func GetTail(head *ListNode) (tail *ListNode) {
	p := head
	for p.Next != nil {
		p = p.Next
	}

	tail = p
	return
}

func main() {

	/*
		head := mockLinkList(1, 2, 3, 4, 5)
		//	PrintList(head)
		PrintE(head, 2)
		rl := reverseKGroup(head, 2)
		PrintE(rl, 2)

		data2 := mockLinkList(1, 2, 3, 4)
		PrintE(data2, 3)
		p := reverseKGroup(data2, 3)
		PrintE(p, 3)
	*/
	k := 2
	data3 := mockLinkList(1, 2, 3, 4, 5, 6, 7)
	PrintE(data3, k)
	res := reverseKGroup(data3, k)
	PrintE(res, k)

}
