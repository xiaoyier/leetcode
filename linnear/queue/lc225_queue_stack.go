package queue


// 请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通队列的全部四种操作（push、top、pop 和 empty）。

// https://leetcode-cn.com/problems/implement-stack-using-queues/

type MyStack struct {
	inQueue []int
	outQueue []int
	toped  bool
}


/** Initialize your data structure here. */
func Constructor1() MyStack {
	return MyStack{
		inQueue: make([]int, 0),
		outQueue: make([]int, 0),
	}
}


// 0 1 2 3 4 5

/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	length := len(this.inQueue)
	if length > 0 {
		for i := 0; i < length; i++ {
			this.outQueue = append(this.outQueue, this.inQueue[0])
			this.inQueue = this.inQueue[1:]
		}
	}
	// 确保 in queue 只有最新的元素
	this.inQueue = append(this.inQueue, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.toped && len(this.outQueue) == 1 {
		ele := this.outQueue[0]
		this.outQueue = this.outQueue[1:]
		this.toped = false
		return ele
	}

	if len(this.inQueue) > 0 {
		ele := this.inQueue[0]
		this.inQueue = this.inQueue[1:]
		return ele
	}

	outLen := len(this.outQueue)
	if outLen > 0 {
		for i := 0; i < outLen; i++ {
			if i == outLen - 1 {
				ele := this.outQueue[0]
				this.outQueue = this.outQueue[1:]
				return ele
			}
			this.inQueue = append(this.inQueue, this.outQueue[0])
			this.outQueue = this.outQueue[1:]
		}
	}
	return 0
}


/** Get the top element. */
func (this *MyStack) Top() int {
	if len(this.inQueue) > 0 {
		ele := this.inQueue[0]
		return ele
	}

	outLen := len(this.outQueue)
	if outLen > 0 {
		for i := 0; i < outLen; i++ {
			if i == outLen - 1 {
				this.toped = true
				return this.outQueue[0]
			}
			this.inQueue = append(this.inQueue, this.outQueue[0])
			this.outQueue = this.outQueue[1:]
		}
	}
	return 0
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.inQueue) == 0 && len(this.outQueue) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
