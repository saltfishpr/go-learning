package main

type MyListNode struct {
	key  int
	val  int
	prev *MyListNode
	next *MyListNode
}

type LRUCache struct {
	capacity int
	head     *MyListNode
	tail     *MyListNode
	catalog  map[int]*MyListNode
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{
		capacity: capacity,
		head:     &MyListNode{},
		tail:     &MyListNode{},
		catalog:  map[int]*MyListNode{},
	}
	c.head.next = c.tail
	c.tail.prev = c.head
	return c
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.catalog[key]
	if ok {
		this.remove(node)
		this.insert(0, node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.catalog[key]
	if ok {
		node.val = value
		this.remove(node)
		this.insert(0, node)
	} else {
		node = &MyListNode{
			key: key,
			val: value,
		}
		this.catalog[key] = node
		this.insert(0, node)
		if len(this.catalog) > this.capacity {
			expiredNode := this.tail.prev
			delete(this.catalog, expiredNode.key)
			this.remove(expiredNode)
		}
	}
}

func (this *LRUCache) insert(idx int, node *MyListNode) {
	cur := this.head
	for i := 0; i < idx; i++ {
		if cur.next == nil {
			break
		}
		cur = cur.next
	}
	cur.next, cur.next.prev, node.next, node.prev = node, node, cur.next, cur
}

func (this *LRUCache) remove(node *MyListNode) {
	node.prev.next, node.next.prev = node.next, node.prev
}
