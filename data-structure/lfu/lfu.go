// @file: lfu.go
// @date: 2021/2/4

// Package lfu
package lfu

type node struct {
	key  int
	next *node
	prev *node
}

type doublyLinkedList struct {
	size int
	head *node
	tail *node
}

func newDoublyLinkedList() *doublyLinkedList {
	head, tail := &node{key: 0}, &node{key: 0}
	head.next = tail
	tail.prev = head

	return &doublyLinkedList{
		size: 0,
		head: head,
		tail: tail,
	}
}

func (l *doublyLinkedList) append(x *node) {
	x.next = l.tail
	x.prev = l.tail.prev
	l.tail.prev.next = x
	l.tail.prev = x

	l.size++
}

func (l *doublyLinkedList) remove(x *node) {
	x.prev.next = x.next
	x.next.prev = x.prev

	l.size--
}

func (l *doublyLinkedList) removeFirst() *node {
	x := l.head.next
	l.remove(x)
	return x
}

type linkedHashSet struct {
	table map[int]*node
	order *doublyLinkedList
}

func newLinkedHashSet() *linkedHashSet {
	order := newDoublyLinkedList()
	return &linkedHashSet{table: make(map[int]*node, 0), order: order}
}

func (l *linkedHashSet) add(key int) {
	if _, ok := l.table[key]; ok {
		return
	}
	x := &node{key: key}
	l.table[key] = x
	l.order.append(x)
}

func (l *linkedHashSet) remove(key int) {
	if _, ok := l.table[key]; !ok {
		return
	}
	x := l.table[key]
	l.order.remove(x)
	delete(l.table, key)
}

type LFUCache struct {
	keyToVal  map[int]int            // key 到 value 的映射
	keyToFreq map[int]int            // key 到 frequency 的映射
	freqToKey map[int]*linkedHashSet // frequency 到 key 列表的映射，存储时间顺序

	minFreq int // 记录最小的频次
	cap     int // 容量
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		keyToVal:  make(map[int]int, 0),
		keyToFreq: make(map[int]int, 0),
		freqToKey: make(map[int]*linkedHashSet, 0),
		minFreq:   0,
		cap:       capacity,
	}
}

func (c *LFUCache) Get(key int) int {
	if _, ok := c.keyToVal[key]; !ok {
		return -1
	}
	// 增加key对应的freq
	c.increaseFreq(key)
	return c.keyToVal[key]
}

func (c *LFUCache) Put(key int, value int) {
	if c.cap <= 0 {
		return
	}
	// key 存在，修改对应的 val 即可
	if _, ok := c.keyToVal[key]; ok {
		c.keyToVal[key] = value
		c.increaseFreq(key)
		return
	}
	// key 不存在，需要插入 key
	if c.cap <= len(c.keyToVal) {
		c.removeMinFreqKey() // 容量若已满淘汰一个freq最小的key
	}

	c.keyToVal[key] = value
	c.keyToFreq[key] = 1
	if _, ok := c.freqToKey[1]; !ok {
		c.freqToKey[1] = newLinkedHashSet()
	}
	c.freqToKey[1].add(key)
	c.minFreq = 1
}

// increaseFreq 增加 key 的使用频率
func (c *LFUCache) increaseFreq(key int) {
	freq := c.keyToFreq[key]
	c.keyToFreq[key]++
	c.freqToKey[freq].remove(key)
	if _, ok := c.freqToKey[freq+1]; !ok {
		c.freqToKey[freq+1] = newLinkedHashSet()
	}
	c.freqToKey[freq+1].add(key)

	// 如果移除后freq列表空，则移除这个键
	if len(c.freqToKey[freq].table) == 0 {
		delete(c.freqToKey, freq)
		// 如果该 freq 是 minFreq
		if freq == c.minFreq {
			c.minFreq++
		}
	}
}

// removeMinFreqKey 删除使用频率最小且最久未使用的 key
func (c *LFUCache) removeMinFreqKey() {
	keyList := c.freqToKey[c.minFreq]
	x := keyList.order.removeFirst()
	delete(keyList.table, x.key)
	if len(keyList.table) == 0 {
		delete(c.freqToKey, c.minFreq)
	}
	delete(c.keyToVal, x.key)
	delete(c.keyToFreq, x.key)
}
