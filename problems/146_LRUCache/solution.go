package main

type LRUCacheNode struct {
	Key        int
	Value      int
	Prev, Next *LRUCacheNode
}

type LRUCache struct {
	Capacity   int
	CacheMap   map[int]*LRUCacheNode
	Head, Tail *LRUCacheNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		CacheMap: make(map[int]*LRUCacheNode, capacity),
		Head:     nil,
		Tail:     nil,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.CacheMap[key]; exists {
		this.moveToHead(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.CacheMap[key]; exists {
		node.Value = value
		this.moveToHead(node)
		return
	}

	if len(this.CacheMap) == this.Capacity {
		this.removeTail()
	}

	newNode := &LRUCacheNode{
		Key:   key,
		Value: value,
	}

	this.addToHead(newNode)
	this.CacheMap[key] = newNode
}

func (this *LRUCache) addToHead(node *LRUCacheNode) {
	node.Prev = nil
	node.Next = this.Head

	if this.Head != nil {
		this.Head.Prev = node
	}

	this.Head = node

	if this.Tail == nil {
		this.Tail = node
	}
}

func (this *LRUCache) moveToHead(node *LRUCacheNode) {
	if this.Head == node {
		return
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if node == this.Tail {
		this.Tail = node.Prev
	}

	this.addToHead(node)
}

func (this *LRUCache) removeTail() {
	if this.Tail == nil {
		return
	}

	delete(this.CacheMap, this.Tail.Key)

	if this.Tail.Prev != nil {
		this.Tail.Prev.Next = nil
	} else {
		this.Head = nil
	}

	this.Tail = this.Tail.Prev
}

// Head -> [A] <-> [B] <-> [C] <- Tail

func main() {
	// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
	// [[2], [1,1], [2,2], [1], [3,3], [2], [4,4], [1], [3], [4]]
	// Expected: [null, null, null, 1, null, -1, null, -1, 3, 4]

	cache := Constructor(2) // null

	cache.Put(1, 1)       // null
	cache.Put(2, 2)       // null
	println(cache.Get(1)) // 1

	cache.Put(3, 3)       // null (evicts key 2)
	println(cache.Get(2)) // -1 (not found)

	cache.Put(4, 4)       // null (evicts key 1)
	println(cache.Get(1)) // -1 (not found)
	println(cache.Get(3)) // 3
	println(cache.Get(4)) // 4
}
