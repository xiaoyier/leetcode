package queue

import (
	"fmt"
	"testing"
)

func TestQueueStack(t *testing.T) {

	stack := Constructor1()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Top())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Top())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Top())
	fmt.Println(stack.Empty())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Empty())
}
