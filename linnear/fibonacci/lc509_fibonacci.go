package fibonacci

// 写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：
// F(0) = 0,   F(1) = 1
// F(N) = F(N - 1) + F(N - 2), 其中 N > 1.

// https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/

// 0 1 1 2 3 5 8 13 21
// 0 1 2 3 4 5 6 7  8
// 时间复杂度 O(n), 空间复杂度 O(1)
func fib101(n int64) int64 {
	if n <= 1 {return n}

	var v1, v2  int64 = 0, 1
	for n > 1 {
		v2, v1 = v2 + v1, v2
		v2 %= 1000000007
		//if v2 >= 1000000007 {
		//	v2 -= 1000000007
		//}
		n--
	}

	//return v2 % 1000000007
	return v2
}

// https://leetcode-cn.com/problems/fibonacci-number/
// 时间复杂度 O(n), 空间复杂度 O(1)
func fib509(n int) int {
	if n <= 1 {return n}

	var v1, v2  int = 0, 1
	for n > 1 {
		v2, v1 = v2 + v1, v2
		n--
	}

	return v2
}

// 时间复杂度 O(n^2), 空间复杂度 O(n)
func fib1(n int) int {
	if n <= 1 {return n}

	return fib1(n-1) + fib1(n-2)
}