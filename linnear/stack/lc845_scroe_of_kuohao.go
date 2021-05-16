package stack

// 给定一个平衡括号字符串 S，按下述规则计算该字符串的分数：
// () 得 1 分。
// AB 得 A + B 分，其中 A 和 B 是平衡括号字符串。
// (A) 得 2 * A 分，其中 A 是平衡括号字符串。

// () -> 1 (()) -> 2 ()()-> 2 ()()() -> 3 (()(())) -> 6   (()())() -> 5
func scoreOfParentheses(s string) int {

	// ((()))((()()))  -> 12

	// 0
	// 0 0 0
	// 0 1
	// 2
	// []  score += 4
	// 0
	// 0 0
	// 0 0 0
	// 0 1
	// 0 1 0
	// 0 2
	// 4
	// [] score += 8

	stack := NewStack()
	score := 0
	for _, char := range []byte(s) {
		if char ==  '(' {
			stack.push(0)
		} else {
			val := stack.pop().(int)
			if val == 0 {
				val = 1
			} else {
				val *= 2
			}
			if stack.isEmpty() {
				score += val
			} else {
				top := stack.pop().(int)
				top += val
				stack.push(top)
			}
		}
	}

	return score
}

// ((()))((()()))  -> 12
func scoreOfParenthese2(s string) int {

	score := 0
	depth := 0
	for i, char := range []byte(s) {
		if char == '(' {
			depth++
		}else {
			depth--
			if []byte(s)[i-1] == '(' {
				score += 1 << depth
			}
		}
	}
	return score
}
