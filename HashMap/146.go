package HashMap

/**
@author Shitanford
@create 2021-09-28 11:38
*/

func main() {

}

type LinkedNode struct {
	prev, next *LinkedNode
	value int
	key int
}

type LRUCache struct {
	capacity int         //capacity of cache
	evict    *LinkedNode //if cache is full, which will be evicted
	recently *LinkedNode //recently used
	cache    map[int]LinkedNode
}

func Constructor(capacity int) LRUCache {
	sentinel := LinkedNode{}	//it is important
	return LRUCache{
		capacity:  capacity,
		evict:     &sentinel,
		recently:  &sentinel,
		cache:     make(map[int]LinkedNode, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.cache[key]; exists {
		//update used linked list
		this.getUpdate(node)
		return node.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int)  {
	if len(this.cache) == this.capacity {
		evictKey := this.evict.key
		this.evict = this.evict.next
		delete(this.cache, evictKey)
	}
	node := LinkedNode{
		prev: this.recently,
		next: nil,
		value: value,
	}
	this.cache[key] = node
	this.recently.next = &node
	this.recently = &node
}

func (this *LRUCache) getUpdate(node LinkedNode)  {
	node.prev.next = node. next
	node.next.prev = node.prev
	node.prev = this.recently
	node.next = nil
	this.recently.next = &node
}
