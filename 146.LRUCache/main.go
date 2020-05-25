package main

type LRUCache struct {
	Size       int
	Capactiy   int
	Cache      map[int]*LinkedNode
	head, tail *LinkedNode
}

type LinkedNode struct {
	key, value int
	prev, next *LinkedNode
}

func initLinkedNode(key, value int) *LinkedNode {
	return &LinkedNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	LRU := LRUCache{
		Capactiy: capacity,
		Cache:    map[int]*LinkedNode{},
		head:     initLinkedNode(0, 0),
		tail:     initLinkedNode(0, 0),
	}
	LRU.head.next = LRU.tail
	LRU.tail.prev = LRU.head
	return LRU
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.Cache[key]; !ok {
		return -1
	}

	node := this.Cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.Cache[key]; !ok {
		node := initLinkedNode(key, value)
		this.Cache[key] = node
		this.addToHead(node)
		this.Size++
		if this.Size > this.Capactiy {
			remove := this.removeTail()
			delete(this.Cache, remove.key)
			this.Size--
		}
	} else {
		node := this.Cache[key]
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
}

func (this *LRUCache) moveToHead(node *LinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *LinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func main() {

}
