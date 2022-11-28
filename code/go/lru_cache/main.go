package main

import "fmt"

type DLinkedNode struct {
	Prev *DLinkedNode
	Next *DLinkedNode
	Key  int
	Val  int
}

type LRUCache struct {
	head, tail *DLinkedNode
	size       int
	capacity   int // 队列最大长度
	cache      map[int]*DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode),
	}

	l.head.Next = l.tail
	l.tail.Prev = l.head
	return l
}

func initDLinkedNode(key, val int) *DLinkedNode {
	return &DLinkedNode{
		Key: key,
		Val: val,
	}
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
	return
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev = prev
	return
}

// 链表都是带虚拟头节点的
func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.Prev = this.head
	node.Next = this.head.Next
	this.head.Next.Prev = node
	this.head.Next = node
	return
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.Prev
	this.removeNode(node)
	return node
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}

	node := this.cache[key]
	this.moveToHead(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.Val = value
		this.moveToHead(node)
		return
	}

	if this.size >= this.capacity {
		removed := this.removeTail()
		delete(this.cache, removed.Key)
		this.size--
	}
	node := &DLinkedNode{Key: key, Val: value}
	this.addToHead(node)
	this.cache[key] = node
	this.size++
	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func Println(cache LRUCache) {
	p := cache.head
	for p != nil {
		fmt.Printf("p.Key=%d, p.Val=%d, p=%p, p.Next=%p, p.Prev=%p\n", p.Key, p.Val, p, p.Next, p.Prev)
		p = p.Next
	}
}
func main() {

	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	Println(cache)
	fmt.Println("-------")
	cache.Get(2)
	Println(cache)

}
