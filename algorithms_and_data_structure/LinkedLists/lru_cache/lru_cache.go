package main

type DoublyLinkedListNode struct {
	key  int
	val  int
	prev *DoublyLinkedListNode
	next *DoublyLinkedListNode
}

type LRUCache struct {
	capacity int
	hashmap  map[int]*DoublyLinkedListNode
	head     *DoublyLinkedListNode
	tail     *DoublyLinkedListNode
}

func Constructor(capacity int) LRUCache {
	head := &DoublyLinkedListNode{key: -1, val: -1}
	tail := &DoublyLinkedListNode{key: -1, val: -1}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		hashmap:  make(map[int]*DoublyLinkedListNode),
		head:     head,
		tail:     tail,
	}
}

func (lru *LRUCache) Get(key int) int {
	node, exists := lru.hashmap[key]
	if !exists {
		return -1
	}
	lru.removeNode(node)
	lru.addToTail(node)
	return node.val
}

func (lru *LRUCache) Put(key int, value int) {
	if node, exists := lru.hashmap[key]; exists {
		lru.removeNode(node)
	}
	node := &DoublyLinkedListNode{key: key, val: value}
	lru.hashmap[key] = node

	if len(lru.hashmap) > lru.capacity {
		lruHead := lru.head.next
		delete(lru.hashmap, lruHead.key)
		lru.removeNode(lruHead)
	}
	lru.addToTail(node)
}

func (lru *LRUCache) addToTail(node *DoublyLinkedListNode) {
	prevNode := lru.tail.prev
	node.prev = prevNode
	node.next = lru.tail
	prevNode.next = node
	lru.tail.prev = node
}

func (lru *LRUCache) removeNode(node *DoublyLinkedListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
