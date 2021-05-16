package stack

import "strconv"

// 根据 逆波兰表示法，求表达式的值。
// 有效的算符包括 +、-、*、/ 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
// 说明：
// 整数除法只保留整数部分。
// 给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。

// https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/

// 2 1 + 3 *  -> (2 + 1) * 3
// 4 13 5 / + -> 13 / 5 + 4

func evalRPN(tokens []string) int {

	stack := NewStack()
	for _, token := range tokens {

		if token == "+" || token == "-" || token == "*" || token == "/" {
			top := stack.pop().(int)
			prev := stack.pop().(int)
			switch token {
			case "+":
				stack.push(prev + top)
			case "-":
				stack.push(prev - top)
			case "*":
				stack.push(prev * top)
			case "/":
				stack.push(prev/top)
			}
		} else {
			num, _ := strconv.Atoi(token)
			stack.push(num)
		}
	}

	return  stack.pop().(int)
}

