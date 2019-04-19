import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

type MinStack struct {
	stack    *ListNode
	minStack *ListNode
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    nil,
		minStack: nil,
	}
}

func (this *MinStack) Push(x int) {
	t := &ListNode{
		Val:  x,
		Next: this.stack,
	}
	this.stack = t
	this.pushMin(x)
}

func (this *MinStack) Pop() {
	if this.stack == nil {
		return
	}

	t := this.stack
	this.stack = this.stack.Next
	t.Next = nil
	this.popMin()
}

func (this *MinStack) Top() int {
	if this.stack == nil {
		return -1
	}
	return this.stack.Val
}

func (this *MinStack) GetMin() int {
	if this.minStack == nil {
		return int(math.Inf(0))
	}

	return this.minStack.Val
}

func (this *MinStack) popMin() {
	if this.minStack == nil {
		return
	}

	t := this.minStack
	this.minStack = this.minStack.Next
	t.Next = nil
}

func (this *MinStack) pushMin(x int) {
	currentMin := this.GetMin()
	if this.minStack == nil {
		currentMin = x
	} else if currentMin > x {
		currentMin = x
	}

	t := &ListNode{
		Val:  currentMin,
		Next: this.minStack,
	}

	this.minStack = t
}
