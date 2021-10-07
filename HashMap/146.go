package main

import "fmt"

/**
@author Shitanford
@create 2021-09-28 11:38
*/

/**
hashMap + linkedList，very good
linkedList can add and remove head and tail element quickly,
but handling middle element is slow, hashMap resolve this case.
*/

func main() {
	//["LRUCache","get","put","get","put","put","get","get"]
	//[[2],[2],[2,6],[1],[1,5],[1,2],[1],[2]]
	lRUCache := Constructor(2)

	lRUCache.Get(2)    // return 1
	lRUCache.Put(2, 6) // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	lRUCache.Get(1)    // returns -1 (not found)
	lRUCache.Put(1, 5) // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	lRUCache.Put(1, 2) // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	lRUCache.Get(1)    // return -1 (not found)
	fmt.Println(lRUCache.Get(2))    // return 3
}

type LinkedNode struct {
	prev, next *LinkedNode
	value int
	key int
}

type LRUCache struct {
	capacity int         //capacity of cache
	sentinel *LinkedNode  //if cache is full, which will be evicted
	cache    map[int]*LinkedNode
}

func Constructor(capacity int) LRUCache {
	node := LinkedNode{}
	node.prev = &node
	node.next = &node
	return LRUCache {
		capacity:  capacity,
		sentinel:  &node,
		cache:     make(map[int]*LinkedNode, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.cache[key]; exists {
		//delete node
		this.deleteNode(node)
		//add node in recently location
		this.addNode(node)
		return node.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int)  {
	if node, exists := this.cache[key]; exists {
		//if key exists, then update value and used list
		node.value = value
		this.deleteNode(node)
		this.addNode(node)
		return
	}
	if len(this.cache) == this.capacity {
		evict := this.sentinel.prev
		this.deleteNode(evict)
		delete(this.cache, evict.key)
	}
	node := LinkedNode{
		value: value,
		key: key,
	}
	this.cache[key] = &node
	//add node
	this.addNode(&node)
}

func (this *LRUCache) deleteNode(node *LinkedNode)  {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) addNode(node *LinkedNode)  {
	node.prev = this.sentinel
	node.next = this.sentinel.next
	this.sentinel.next.prev = node
	this.sentinel.next = node
}