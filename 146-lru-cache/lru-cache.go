type ListNode struct {
	Key  int
	Val  int
	Next *ListNode
	Pre  *ListNode
}

type LRUCache struct {
	Head *ListNode
	Tail *ListNode
	Cap  int
	Len  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Head: nil,
		Tail: nil,
		Cap:  capacity,
		Len:  0,
	}
}

func (cache *LRUCache) Get(key int) int {
	currentNode := cache.findNode(key)
	if currentNode == nil {
		return -1
	}

	cache.moveToFirst(currentNode)
	return currentNode.Val
}

func (cache *LRUCache) Put(key int, value int) {
	currentNode := cache.findNode(key)
	if currentNode == nil {
		if cache.Cap == cache.Len {
			cache.removeLast()
		}

		currentNode = &ListNode{
			Key:  key,
			Val:  value,
			Next: nil,
			Pre:  nil,
		}
		if cache.Head == nil {
			cache.Tail = currentNode
		} else {
			cache.Head.Pre = currentNode
		}
		currentNode.Next = cache.Head
		cache.Head = currentNode
		cache.Len += 1
	} else {
		currentNode.Val = value
		cache.moveToFirst(currentNode)
	}
}

func (cache *LRUCache) moveToFirst(node *ListNode) {
	//node.PrePre为nil 说明是第一个节点,不需要移动
	if node.Pre != nil {
		node.Pre.Next = node.Next
		//最后一个节点的前去不需要设置
		if node.Next != nil {
			node.Next.Pre = node.Pre
		} else {
			cache.Tail = cache.Tail.Pre
		}
		node.Next = cache.Head
		node.Pre = nil
		cache.Head.Pre = node
		cache.Head = node
	}
}

func (cache *LRUCache) findNode(key int) *ListNode {
	for head := cache.Head; head != nil; head = head.Next {
		if head.Key == key {
			return head
		}
	}

	return nil
}

func (cache *LRUCache) removeLast() {
	if cache.Tail == nil {
		return
	}

	// 无前驱，第一个节点
	if cache.Tail.Pre == nil {
		cache.Head = nil
		cache.Tail = nil
	} else {
		cache.Tail = cache.Tail.Pre
		cache.Tail.Next.Pre = nil
		cache.Tail.Next = nil
	}

	cache.Len -= 1
}