package stack

import "fmt"

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 		左括号必须用相同类型的右括号闭合。
//      左括号必须以正确的顺序闭合。

// https://leetcode-cn.com/problems/valid-parentheses/

type Stack struct {
	size int
	elements []interface{}
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]interface{}, 10),
	}
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) isEmpty() bool {
	return s.size == 0
}

func (s *Stack) push(element interface{}) {
	if s.size == len(s.elements) {
		newElements := make([]interface{}, s.size << 1)
		for index, ele := range s.elements {
			newElements[index] = ele
		}
		s.elements = newElements
	}
	s.elements[s.size] = element
	s.size++
}

func (s *Stack) pop() interface{} {
	if s.size == 0 {
		return nil
	}
	element := s.elements[s.size-1]
	s.elements[s.size-1] = nil
	s.size--
	return element
}

// {[()]} -> true
// ()[]{} -> true
func isValid(s string) bool {

	stack := NewStack()
	for _, char := range []byte(s) {
		if char == '{' || char == '[' || char == '(' {
			stack.push(char)
		} else {
			if stack.isEmpty() {
				return false
			}
			c := stack.pop().(byte)
			if char == '}' && c != '{' {
				return false
			}
			if char == ']' && c != '[' {
				return false
			}
			if char == ')' && c != '(' {
				fmt.Printf("%c <---> %c", c, char)
				return false
			}
		}
	}

	for !stack.isEmpty() {
		return false
	}
	return true
}

