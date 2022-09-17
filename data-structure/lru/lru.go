// @file: lru.go
// @date: 2021/2/3

// Package lru
package lru

type node struct {
	key  int
	val  int
	next *node
	prev *node
}

func newNode(key, val int) *node {
	return &node{key: key, val: val}
}

type doubleList struct {
	head *node
	tail *node
	size int
}

func newDoubleList() *doubleList {
	head := newNode(0, 0)
	tail := newNode(0, 0)
	head.next = tail
	tail.prev = head
	return &doubleList{head: head, tail: tail, size: 0}
}

// appendNode add node x at the end of the linked list.
func (l *doubleList) appendNode(x *node) {
	x.prev = l.tail.prev
	x.next = l.tail
	l.tail.prev.next = x
	l.tail.prev = x

	l.size++
}

// removeNode remove node x. (x must exist)
func (l *doubleList) removeNode(x *node) {
	x.prev.next = x.next
	x.next.prev = x.prev

	l.size--
}

// removeFirst delete the first node in the linked list and return it.
func (l *doubleList) removeFirst() *node {
	if l.head.next == l.tail {
		return nil
	}
	x := l.head.next
	l.removeNode(x)
	return x
}

/*
LRU缓存，在双向链表头部的是最久未使用的
*/
type LRUCache struct {
	m     map[int]*node
	cache *doubleList
	cap   int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m:     make(map[int]*node, 0),
		cache: newDoubleList(),
		cap:   capacity,
	}
}

// makeRecently promote the node corresponding to the key to the most recently used.
func (c *LRUCache) makeRecently(key int) {
	x := c.m[key]
	c.cache.removeNode(x)
	c.cache.appendNode(x)
}

// addRecently add a node to cache.
func (c *LRUCache) addRecently(key, value int) {
	x := &node{key: key, val: value}
	c.cache.appendNode(x)
	c.m[key] = x
}

// deleteKey delete a key.
func (c *LRUCache) deleteKey(key int) {
	x := c.m[key]
	c.cache.removeNode(x)
	delete(c.m, key)
}

// removeLeastRecently delete the least used element.
func (c *LRUCache) removeLeastRecently() {
	x := c.cache.removeFirst()
	delete(c.m, x.key)
}

// Get read data from cache.
func (c *LRUCache) Get(key int) int {
	if _, ok := c.m[key]; !ok {
		return -1
	}
	c.makeRecently(key)
	return c.m[key].val
}

// Put insert key and val into cache.
func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.m[key]; ok {
		c.deleteKey(key)
		c.addRecently(key, value)
		return
	}
	if c.cache.size == c.cap {
		c.removeLeastRecently()
	}
	c.addRecently(key, value)
}
