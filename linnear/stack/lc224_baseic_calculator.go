package stack


// 给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。

// https://leetcode-cn.com/problems/basic-calculator/

// 1 + 1 -> 2
// 1 - (1 - 3) -> 3

func calculate(s string) int {
	// 记录当前正负数
	stack := []int{1}
	result := 0
	num := 0
	op := 1
	for _, c := range []byte(s) {

		if c >= '0' && c <= '9' {
			num = num * 10 + int(c-'0')
			continue
		}

		result += num * op
		num = 0

		switch c {
		case ' ': // do nothing
		case '+':
			op = stack[len(stack)-1]
		case '-':
			op = -stack[len(stack)-1]
		case '(':
			stack = append(stack, op)
		case ')':
			stack = stack[:len(stack)-1]
		}
	}

	result += num * op

	return  result
}