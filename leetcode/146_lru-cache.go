package main

// https://leetcode.cn/problems/lru-cache/description/

type LRUCache struct {
	size, capacity int
	cache          map[int]*LinkedNode
	head, tail     *LinkedNode
}

type LinkedNode struct {
	key, value int
	prev, next *LinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    make(map[int]*LinkedNode),
		head:     &LinkedNode{},
		tail:     &LinkedNode{},
	}

	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}

	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; !ok {
		n := &LinkedNode{
			key:   key,
			value: value,
		}
		this.cache[key] = n
		this.addToHead(n)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *LinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *LinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil
}

func (this *LRUCache) removeTail() *LinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func (this *LRUCache) moveToHead(node *LinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
