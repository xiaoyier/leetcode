package queue

// 请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：

// https://leetcode-cn.com/problems/implement-queue-using-stacks/


type MyQueue struct {
	inStack []int
	outStack []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		inStack: make([]int, 0),
		outStack: make([]int, 0),
	}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.inStack = append(this.inStack, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	ele := this.Peek()
	this.outStack = this.outStack[:len(this.outStack) - 1]
	return ele
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		for i := 0; i < len(this.inStack); i++ {
			ele := this.inStack[len(this.inStack)-1 - i]
			this.outStack = append(this.outStack, ele)
		}
		this.inStack = make([]int, 0)
	}

	ele := this.outStack[len(this.outStack) - 1]
	return ele
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
