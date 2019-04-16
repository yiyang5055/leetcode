type ListNode struct {
	Key  int
	Val  int
	Next *ListNode
}

type LRUCache struct {
	Head *ListNode
	Cap  int
	Len  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Head: nil,
		Cap:  capacity,
		Len:  0,
	}
}

func (cache *LRUCache) Get(key int) int {
	if cache.Len < 0 {
		return -1
	}

	preNode, currentNode := cache.findNode(key)
	if currentNode != nil {
		// preNode为nil 说明只有一个节点
		if preNode != nil {
			preNode.Next = currentNode.Next
			currentNode.Next = cache.Head
			cache.Head = currentNode
		}

		return cache.Head.Val
	}
	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	preNode, currentNode := cache.findNode(key)
	if currentNode == nil {
		if cache.Cap == cache.Len {
			cache.removeLast()
		}

		currentNode = &ListNode{
			Key:  key,
			Val:  value,
			Next: nil,
		}
		currentNode.Next = cache.Head
		cache.Head = currentNode
		cache.Len += 1
	} else {
		currentNode.Val = value
		if preNode != nil {
			preNode.Next = currentNode.Next
			currentNode.Next = cache.Head
			cache.Head = currentNode
		}
	}
}

func (cache *LRUCache) findNode(key int) (*ListNode, *ListNode) {
	preNode := (*ListNode)(nil)
	currentNode := cache.Head
	for currentNode != nil {
		if currentNode.Key == key {
			return preNode, currentNode
		}
		preNode = currentNode
		currentNode = currentNode.Next
	}
	return nil, nil
}

func (cache *LRUCache) removeLast() {
	if cache.Len < 2 {
		cache.Head = nil
		cache.Len -= 1
		return
	}

	preNode := cache.Head
	behindNode := preNode.Next
	for behindNode.Next != nil {
		preNode = behindNode
		behindNode = behindNode.Next
	}
	preNode.Next = nil
	cache.Len -= 1
	return
}